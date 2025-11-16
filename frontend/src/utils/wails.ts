import { throttle } from 'radash';
import { WailsService } from '../../bindings/app.local/app';
import { WindowStatus, ActionOpt } from '../../bindings/app.local/app/types';

import { BallWindowType } from '../../bindings/app.local/app/global';

import * as runtime from '@wailsio/runtime';

export { WindowStatus, ActionOpt, BallWindowType };
export const WailsRuntime = runtime;
export const WailsServe = WailsService;

// 节流函数，防止短时间内多次点击打开多个窗口
export const OpenTestWindow = throttle({ interval: 1000 }, () => {
  WailsService.OpenTestWindow();
});

export const OpenBallWindow = throttle({ interval: 1000 }, () => {
  WailsService.OpenBallWindow();
});
