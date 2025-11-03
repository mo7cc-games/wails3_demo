#!/usr/bin/env ts-node
/* 
清理目录 ，这个脚本会直接删除某些目录
*/
import path from 'path';
import fs from 'fs-extra';

const clearPath = [
  'node_modules',
  '.cache',
  'dist',
  'bun.lock',
  'go.sum',
  'docs/.vitepress/cache',
  'docs/.vitepress/dist',
  '.trash',
  'package-lock.json',
  // ...
];

// 获得当前目录的绝对路径
const currentDir = process.cwd();

// 遍历删除目录，如果目录不存在则跳过
for (const relPath of clearPath) {
  const fullPath = path.join(currentDir, relPath);
  if (fs.existsSync(fullPath)) {
    console.info(`Removing: ${fullPath}`);
    fs.removeSync(fullPath);
  }
}

// 进程 结束
process.exit(0);
