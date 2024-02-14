// frontend/main.js

// 引入Code Notation Manager框架
import { CNM } from './cnm/cnm.js';

// 主要逻辑入口
function main() {
  // 初始化Code Notation Manager
  const cnm = new CNM();

  // 添加示例模块
  cnm.addModule('feature1', feature1Module);
  cnm.addModule('feature2', feature2Module);

  // 启动应用
  cnm.start();
}

// 示例模块1
function feature1Module() {
  console.log('Feature 1 Module Loaded');

  // 在这里添加Feature 1的具体逻辑
}

// 示例模块2
function feature2Module() {
  console.log('Feature 2 Module Loaded');

  // 在这里添加Feature 2的具体逻辑
}

// 在页面加载完成后运行主要逻辑
document.addEventListener('DOMContentLoaded', main);
