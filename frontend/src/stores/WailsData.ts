import { defineStore } from 'pinia';
import { useCounterOptionsStore } from './counterOptions';
import {
  WailsRuntime,
  WailsServe,
  WindowStatus,
  ActionOpt,
  BallWindowType,
} from '@src/utils/wails';
import { debounce } from 'radash';
import { OpenBallWindow } from '@src/utils/wails';

type WailsDataStateType = {
  TimeStr: string;
  WindowName: string;
  IsWebGPU: boolean;
  IsFrameless: boolean;
  ScreenInfo: WindowStatus;
  BallWindow: BallWindowType;
};

type ActionVal = {
  name: string;
  data: ActionOpt;
};

export const useWailsDataStore = defineStore('WailsData', {
  state(): WailsDataStateType {
    return {
      WindowName: '',
      IsWebGPU: false,
      TimeStr: '',
      IsFrameless: false,
      ScreenInfo: {} as WindowStatus,
      BallWindow: {} as BallWindowType,
    };
  },
  actions: {
    SetFrameless() {
      const _this = this;
      WailsRuntime.Window.SetFrameless(true).then(() => {
        _this.IsFrameless = true;
      });
    },
    OutFrameless() {
      const _this = this;
      WailsRuntime.Window.SetFrameless(false).then(() => {
        _this.IsFrameless = false;
      });
    },
  },
});

const _GetScreenInfo = async () => {
  const WailsDataStore = useWailsDataStore();
  if (!WailsDataStore.WindowName) {
    return;
  }
  let IsExtended = false;
  // @ts-expect-error: 如果不存在该属性就 false
  if (window.screen.isExtended) {
    IsExtended = true;
  }
  const ScreenInfo: WindowStatus = {
    WindowName: WailsDataStore.WindowName, // 窗口名称
    Width: WailsDataStore.ScreenInfo.Width, // 屏幕宽度
    Height: WailsDataStore.ScreenInfo.Height, // 屏幕高度
    AvailWidth: window.screen.availWidth, // 可用屏幕宽度
    AvailHeight: window.screen.availHeight, // 可用屏幕高度
    // 如果为 0（例如最左/最顶位置），使用 nullish 合并以保留 0 值
    ScreenLeft: WailsDataStore.ScreenInfo.ScreenLeft, // 距离屏幕左侧位置
    ScreenTop: WailsDataStore.ScreenInfo.ScreenTop, // 距离屏幕顶部位置
    // 正确映射内部宽高（之前被误写反了）
    InnerWidth: window.innerWidth, // 内容区域宽度 包括滚动条
    InnerHeight: window.innerHeight, // 内容区域高度
    Dpr: window.devicePixelRatio, // 当前 DPR
    IsFullscreen: WailsDataStore.ScreenInfo.IsFullscreen, // 是否全屏
    IsExtended, // 是否为扩展屏幕
    IsMaximise: WailsDataStore.ScreenInfo.IsMaximise, // 是否最大化
    IsMinimise: WailsDataStore.ScreenInfo.IsMinimise, // 是否最小化
    IsFocused: WailsDataStore.ScreenInfo.IsFocused, // 是否为焦点窗口
  };

  ScreenInfo.IsFullscreen = await WailsRuntime.Window.IsFullscreen();
  ScreenInfo.IsMaximise = await WailsRuntime.Window.IsMinimised();
  ScreenInfo.IsMaximise = await WailsRuntime.Window.IsMaximised();
  const Pos = await WailsRuntime.Window.Position();
  ScreenInfo.ScreenLeft = Pos.x;
  ScreenInfo.ScreenTop = Pos.y;

  const screen = await WailsRuntime.Window.GetScreen();
  ScreenInfo.Width = screen.Size.Width;
  ScreenInfo.Height = screen.Size.Height;

  // console.log('ScreenInfo', JSON.stringify(ScreenInfo, null, 2));
  // 传递给 Golang 端
  WailsServe.WindowChange(ScreenInfo);
  WailsDataStore.ScreenInfo = ScreenInfo;
};

// 获取屏幕信息 - 窗口移动+大小变化时调用 - 防抖处理
const GetScreenInfo = debounce({ delay: 100 }, () => {
  _GetScreenInfo();
});

const GetAllWindowInfo = debounce({ delay: 100 }, () => {
  // 获取所有窗口的状态和信息
  // WailsServe.GetAllWindowInfo().then((allWindow) => {
  //   console.info('allWindow', allWindow);
  // });
});

export const StartWailsDataListener = async () => {
  // 等待资源全部加载完毕 再去干一些事情
  const WailsDataStore = useWailsDataStore();
  WailsRuntime.Window.Name().then((name) => {
    WailsDataStore.WindowName = name;
    if (window.navigator.gpu) {
      WailsDataStore.IsWebGPU = true;
    }
    GetScreenInfo();
    GetAllWindowInfo();

    // 默认打开悬浮球窗口
    OpenBallWindow();

    WailsServe.GetWindowInfo(name).then((info) => {
      if (info.EnableFrameless) {
        WailsDataStore.SetFrameless();
      } else {
        WailsDataStore.OutFrameless();
      }
    });

    // 初始化获取悬浮球信息
    WailsServe.GetBallWindowZoomDir().then((dir) => {
      WailsDataStore.BallWindow = dir;
    });
  });

  // 来自 Wails 的事件监听
  WailsRuntime.Events.On('Time', (val) => {
    if (val.data) {
      const TimeStr = val.data;
      WailsDataStore.TimeStr = TimeStr;
    }
  });

  // 专门用于 Pinia 的 Action 通信，可以直接使得多窗口的数据保持一致性
  // 用法 WailsServe.Action('counterOptions.add');
  WailsRuntime.Events.On('Action', (val: ActionVal) => {
    if (val.data) {
      const data = val.data;
      if (data.ActionName === 'counterOptions.add') {
        const store = useCounterOptionsStore();
        store.add();
      }
    }
  });

  WailsRuntime.Events.On('WindowChange', (val: ActionVal) => {
    GetAllWindowInfo();
    if (val.data) {
      const data = val.data;
      if (data.WindowName !== WailsDataStore.WindowName) {
        return;
      }
      GetScreenInfo();
    }
  });

  WailsRuntime.Events.On('BallWindowZoom', (val) => {
    if (val.data) {
      WailsDataStore.BallWindow = val.data;
    }
  });

  // 监听窗口失焦
  window.addEventListener('blur', () => {
    WailsDataStore.ScreenInfo.IsFocused = false;
    _GetScreenInfo();
  });

  // 监听窗口聚焦
  window.addEventListener('focus', () => {
    WailsDataStore.ScreenInfo.IsFocused = true;
    _GetScreenInfo();
  });
};
