// frontend/cnm/cnm.js

class CNM {
  constructor() {
    this.modules = {};
  }

  addModule(name, moduleFunction) {
    this.modules[name] = moduleFunction;
  }

  start() {
    console.log('Code Notation Manager Started');

    // 启动每个模块
    for (const moduleName in this.modules) {
      if (this.modules.hasOwnProperty(moduleName)) {
        const moduleFunction = this.modules[moduleName];
        moduleFunction();
      }
    }
  }
}

export { CNM };
