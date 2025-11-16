import { fileURLToPath, URL } from 'node:url';
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import vueDevTools from 'vite-plugin-vue-devtools';
import path from 'node:path';
import dotenv from 'dotenv';
import legacy from '@vitejs/plugin-legacy';
import wails from '@wailsio/runtime/plugins/vite';

const __dirname = fileURLToPath(new URL('.', import.meta.url));
console.info('Vue3项目启动,__dirname', __dirname);

const envPath = path.resolve(__dirname, '../.env');
const env = dotenv.config({ path: envPath });

let port = 5173; // 默认端口
if (env.parsed) {
  if (env.parsed.WAILS_VITE_PORT) {
    port = parseInt(env.parsed.WAILS_VITE_PORT, 10);
  }
}
console.info('Vue3设定端口号为', port);

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    wails('./bindings'),
    legacy({
      targets: ['since 2020', 'not dead'],
    }),
    vueDevTools(),
  ],
  resolve: {
    alias: {
      '@src': path.resolve(__dirname, './src'),
    },
  },
  server: {
    host: '0.0.0.0',
    port,
  },
});
