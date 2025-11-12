#!/usr/bin/env ts-node

/* 
将所有依赖全部更到最新，然后删除并重新安装
*/
import shell from 'shelljs';
import path from 'path';

if (!shell.which('go')) {
  shell.echo('Sorry, this script requires go');
  shell.exit(1);
}

if (!shell.which('npm')) {
  shell.echo('Sorry, this script requires npm');
  shell.exit(1);
}

const currentDir = process.cwd();
shell.cd(currentDir);
console.info('当前执行命令的目录:', shell.pwd().toString());

console.info('更新并安装 Wails3 最新版本...');
shell.exec('go install github.com/wailsapp/wails/v3/cmd/wails3@latest');
shell.exec('wails3 version');
shell.exec('go get -u ./...');
shell.rm('-rf', 'go.sum');
shell.exec('go mod tidy');

const frontendPath = path.join(currentDir, 'frontend');
shell.cd(frontendPath);
console.info('进入前端目录:', shell.pwd().toString());
shell.rm('-rf', 'node_modules');
shell.rm('-f', 'package-lock.json');
shell.exec('npm update');

process.exit(0);
