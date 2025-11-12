import { defineStore } from 'pinia';
import wailsio from '@src/utils/wailsio';
import { useCounterOptionsStore } from './counterOptions';
import { WailsService } from '@server/app';
import { WindowChangeOpt } from '@server/app/models';
import { debounce } from 'radash';

type WailsDataState = {
  TimeStr: string;
  WindowName: string;
  IsWebGPU: boolean;
  ScreenInfo: WindowChangeOpt;
};

export const useWailsDataStore = defineStore('WailsData', {
  state(): WailsDataState {
    return {
      WindowName: '',
      IsWebGPU: false,
      TimeStr: '',
      ScreenInfo: {} as WindowChangeOpt,
    };
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
  const ScreenInfo: WindowChangeOpt = {
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
    NowDpr: window.devicePixelRatio, // 当前 DPR
    IsFullscreen: WailsDataStore.ScreenInfo.IsFullscreen, // 是否全屏
    IsExtended, // 是否为扩展屏幕
    IsMaximise: WailsDataStore.ScreenInfo.IsMaximise, // 是否最大化
    IsMinimise: WailsDataStore.ScreenInfo.IsMinimise, // 是否最小化
    IsFocused: WailsDataStore.ScreenInfo.IsFocused, // 是否为焦点窗口
  };

  ScreenInfo.IsFullscreen = await wailsio.Window.IsFullscreen();
  ScreenInfo.IsMaximise = await wailsio.Window.IsMinimised();
  ScreenInfo.IsMaximise = await wailsio.Window.IsMaximised();
  const Pos = await wailsio.Window.Position();
  ScreenInfo.ScreenLeft = Pos.x;
  ScreenInfo.ScreenTop = Pos.y;

  ScreenInfo.IsFocused = await wailsio.Window.IsFocused();

  const wailsScreen = await wailsio.Window.GetScreen();
  ScreenInfo.Width = wailsScreen.Size.Width;
  ScreenInfo.Height = wailsScreen.Size.Height;
  // console.log('ScreenInfo', JSON.stringify(ScreenInfo, null, 2));
  // 传递给 Golang 端
  WailsService.WindowChange(ScreenInfo);
  WailsDataStore.ScreenInfo = ScreenInfo;
};

// 获取屏幕信息 - 窗口移动+大小变化时调用 - 防抖处理
const GetScreenInfo = debounce({ delay: 100 }, () => {
  _GetScreenInfo();
});

type WindowChangeType = {
  Action: string;
  WindowName: string;
};
// 全局只会执行一次的函数
export const StartWailsDataListener = async () => {
  const WailsDataStore = useWailsDataStore();
  // 当前所在窗口的名字，由 Wails 声明
  wailsio.Window.Name().then((name) => {
    WailsDataStore.WindowName = name;
    if (window.navigator.gpu) {
      WailsDataStore.IsWebGPU = true;
    }
    GetScreenInfo();
  });

  // 专门用于 Pinia 的 Action 通信，可以直接使得多窗口的数据保持一致性
  // 用法 WailsService.Action('counterOptions.add');
  wailsio.Events.On('Action', (val) => {
    if (val.data) {
      const name = val.data as string;
      if (name === 'counterOptions.add') {
        const store = useCounterOptionsStore();
        store.add();
      }
    }
  });

  // 来自 Wails 的事件监听
  wailsio.Events.On('Time', (val) => {
    if (val.data) {
      const TimeStr = val.data;
      WailsDataStore.TimeStr = TimeStr;
    }
  });

  wailsio.Events.On('WindowChange', (val) => {
    if (val.data) {
      const data = val.data as WindowChangeType;
      if (data.WindowName !== WailsDataStore.WindowName) {
        return;
      }
      GetScreenInfo();
    }
  });
};
