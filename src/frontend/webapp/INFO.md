# QUASAR PROJECT INFO

Project Config Details

⭕ = todo ✅ = done

---

## create a Quasar app

```bash
yarn create quasar

√ What would you like to build? » App with Quasar CLI, let's go!
√ Project folder: ... quasar-project
√ Pick Quasar version: » Quasar v2 (Vue 3 | latest and greatest)
√ Pick script type: » Javascript
√ Pick Quasar App CLI variant: » Quasar App CLI with Vite 2 (stable | v1)
√ Package name: ... quasar-project
√ Project product name: (must start with letter if building mobile apps) ... Quasar App
√ Project description: ... A Quasar Project
√ Author: ... s.wyss <stefan.arnold.wyss@gmail.com>
√ Pick a Vue component style: » Composition API with <script setup>
√ Pick your CSS preprocessor: » Sass with SCSS syntax
√ Check the features needed for your project: » Linting (vite-plugin-checker + ESLint), State Management (Pinia), axios, vue-i18n
√ Pick an ESLint preset: » Prettier

√ Install project dependencies? (recommended) » Yes, use yarn
```

---

## add modes

### SSR ✅

> https://quasar.dev/quasar-cli-vite/developing-ssr/preparation

```bash
quasar mode add ssr
```

### PWA ✅

> https://quasar.dev/quasar-cli-vite/developing-pwa/preparation

```bash
quasar mode add pwa
```

---

## add packages

### lodash ⭕

Lodash modular utilities.

> https://lodash.com/

```bash
yarn add lodash
```

### dotenv ⭕

Loads environment variables from .env file

> https://github.com/motdotla/dotenv#readme

```bash
yarn add dotenv
```

### moment ⭕

Parse, validate, manipulate, and display dates and times in JavaScript.

> https://momentjs.com/

```bash
yarn add moment
```

### QCalendar ⭕

Build Beautiful, Responsive Calendars

> https://qcalendar.netlify.app/

```bash
quasar ext add @quasar/qcalendar@next
```

### APEXCHARTS.JS ⭕

Modern & Interactive Open-source Charts

> https://apexcharts.com/

```bash
yarn add apexcharts
```

### Chart.js ⭕

Easy and beautiful charts with Chart.js and Vue.js

> https://vue-chartjs.org/

```bash
yarn add vue-chartjs chart.js
yarn add vue3-apexcharts
```

---

## config&design

### Icon Genie CLI ⭕

> https://quasar.dev/icongenie/installation

```bash
# Node.js >=18 is required.
yarn global add @quasar/icongenie
```

#### Icons ⭕

> https://quasar.dev/vue-components/icon > https://iconexplorer.app/?selected
>
> https://icons.getbootstrap.com/

### Layout

> https://quasar.dev/layout-builder

#### MainLayout

#### Separate Layout Components

### Theme Builder ⭕

> https://quasar.dev/style/theme-builder

#### DarkTheme ⭕

> https://quasar.dev/quasar-plugins/dark#dark-api

---

## configure quasar.config.js ℹ️

```js
/* eslint-env node */

/*
 * This file runs in a Node context (it's NOT transpiled by Babel), so use only
 * the ES6 features that are supported by your Node version. https://node.green/
 */

// Configuration for your app
// https://v2.quasar.dev/quasar-cli-vite/quasar-config-js

const {configure} = require("quasar/wrappers");
const path = require("path");

module.exports = configure(function (/* ctx */) {
  return {
    // https://v2.quasar.dev/quasar-cli-vite/prefetch-feature
    // preFetch: true,

    // app boot file (/src/boot)
    // --> boot files are part of "main.js"
    // https://v2.quasar.dev/quasar-cli-vite/boot-files
    boot: ["i18n", "axios"],

    // https://v2.quasar.dev/quasar-cli-vite/quasar-config-js#css
    css: ["app.scss"],

    // https://github.com/quasarframework/quasar/tree/dev/extras
    extras: ["bootstrap-icons"],

    // Full list of options: https://v2.quasar.dev/quasar-cli-vite/quasar-config-js#build
    build: {
      target: {
        browser: ["es2019", "edge88", "firefox78", "chrome87", "safari13.1"], node: "node20",
      },

      vueRouterMode: "hash", // available values: 'hash', 'history'
      // vueRouterBase,
      // vueDevtools,
      // vueOptionsAPI: false,

      // rebuildCache: true, // rebuilds Vite/linter/etc cache on startup

      // publicPath: '/',
      // analyze: true,
      // env: {},
      // rawDefine: {}
      // ignorePublicFolder: true,
      // minify: false,
      // polyfillModulePreload: true,
      // distDir

      // extendViteConf (viteConf) {},
      // viteVuePluginOptions: {},

      vitePlugins: [["@intlify/vite-plugin-vue-i18n", {
        // if you want to use Vue I18n Legacy API, you need to set `compositionOnly: false`
        // compositionOnly: false,

        // if you want to use named tokens in your Vue I18n messages, such as 'Hello {name}',
        // you need to set `runtimeOnly: false`
        // runtimeOnly: false,

        // you need to set i18n resource including paths !
        include: path.resolve(__dirname, "./src/i18n/**"),
      },], ["vite-plugin-checker", {
        eslint: {
          lintCommand: 'eslint "./**/*.{js,mjs,cjs,vue}"',
        },
      }, {server: false},],],
    },

    // Full list of options: https://v2.quasar.dev/quasar-cli-vite/quasar-config-js#devServer
    devServer: {
      // https: true
      open: true, // opens browser window automatically
    },

    // https://v2.quasar.dev/quasar-cli-vite/quasar-config-js#framework
    framework: {
      config: {},

      // iconSet: 'material-icons', // Quasar icon set
      // lang: 'en-US', // Quasar language pack

      // For special cases outside of where the auto-import strategy can have an impact
      // (like functional components as one of the examples),
      // you can manually specify Quasar components/directives to be available everywhere:
      //
      // components: [],
      // directives: [],

      // Quasar plugins
      plugins: [],
    },

    // animations: 'all', // --- includes all animations
    // https://v2.quasar.dev/options/animations
    animations: [],

    // https://v2.quasar.dev/quasar-cli-vite/quasar-config-js#property-sourcefiles
    // sourceFiles: {
    //   rootComponent: 'src/App.vue',
    //   router: 'src/router/index',
    //   store: 'src/store/index',
    //   registerServiceWorker: 'src-pwa/register-service-worker',
    //   serviceWorker: 'src-pwa/custom-service-worker',
    //   pwaManifestFile: 'src-pwa/manifest.json',
    //   electronMain: 'src-electron/electron-main',
    //   electronPreload: 'src-electron/electron-preload'
    // },

    // https://v2.quasar.dev/quasar-cli-vite/developing-ssr/configuring-ssr
    //https://quasar.dev/quasar-cli-webpack/developing-ssr/configuring-ssr/
    ssr: {
      pwa: true, // should a PWA take over (default: false), or just a SPA?

      /**
       * Manually serialize the store state and provide it yourself
       * as window.__INITIAL_STATE__ to the client-side (through a <script> tag)
       * (Requires @quasar/app-webpack v3.5+)
       */
      manualStoreSerialization: false,

      /**
       * Manually inject the store state into ssrContext.state
       * (Requires @quasar/app-webpack v3.5+)
       */
      manualStoreSsrContextInjection: false,

      /**
       * Manually handle the store hydration instead of letting Quasar CLI do it.
       * For Pinia: store.state.value = window.__INITIAL_STATE__
       * For Vuex: store.replaceState(window.__INITIAL_STATE__)
       */
      manualStoreHydration: false,

      /**
       * Manually call $q.onSSRHydrated() instead of letting Quasar CLI do it.
       * This announces that client-side code should takeover.
       */
      manualPostHydrationTrigger: false,

      prodPort: 3000, // The default port that the production server should use
      // (gets superseded if process∙env∙PORT is specified at runtime)

      maxAge: 1000 * 60 * 60 * 24 * 30, // Tell browser when a file from the server should expire from cache
      // (the default value, in ms)
      // Has effect only when server.static() is used

      // List of SSR middleware files (src-ssr/middlewares/*). Order is important.
      middlewares: [// ...
        "render", // Should not be missing, and should be last in the list.
      ],

      // optional; add/remove/change properties
      // of production generated package.json
      extendPackageJson(pkg) {
        // directly change props of pkg;
        // no need to return anything
      },

      // optional;
      // handles the Webserver webpack config ONLY
      // which includes the SSR middleware
      extendWebpackWebserver(cfg) {
        // directly change props of cfg;
        // no need to return anything
      },

      // optional; EQUIVALENT to extendWebpack() but uses webpack-chain;
      // handles the Webserver webpack config ONLY
      // which includes the SSR middleware
      chainWebpackWebserver(chain) {
        // chain is a webpack-chain instance
        // of the Webpack configuration
      },
    },

    // https://v2.quasar.dev/quasar-cli-vite/developing-pwa/configuring-pwa
    pwa: {
      // workboxPluginMode: 'InjectManifest',
      // workboxOptions: {},
      manifest: {
        // ...
      },

      // Use this OR metaVariablesFn, but not both;
      // variables used to inject specific PWA
      // meta tags (below are default values);
      metaVariables: {
        appleMobileWebAppCapable: "yes",
        appleMobileWebAppStatusBarStyle: "default",
        appleTouchIcon120: "icons/apple-icon-120x120.png",
        appleTouchIcon180: "icons/apple-icon-180x180.png",
        appleTouchIcon152: "icons/apple-icon-152x152.png",
        appleTouchIcon167: "icons/apple-icon-167x167.png",
        appleSafariPinnedTab: "icons/safari-pinned-tab.svg",
        msapplicationTileImage: "icons/ms-icon-144x144.png",
        msapplicationTileColor: "#000000",
      },

      /*  // Optional, overrides metaVariables above;
       // Use this OR metaVariables, but not both;
       metaVariablesFn(manifest) {
         // ...
         return [
           {
             // this entry will generate:
             // <meta name="theme-color" content="ff0">

             tagName: "meta",
             attributes: {
               name: "theme-color",
               content: "#ff0",
             },
           },

           {
             // this entry will generate:
             // <link rel="apple-touch-icon" sizes="180x180" href="icons/icon-180.png">
             // references /public/icons/icon-180.png

             tagName: "link",
             attributes: {
               rel: "apple-touch-icon",
               sizes: "180x180",
               href: "icons/icon-180.png",
             },
             closeTag: false, // this is optional;
             // specifies if tag also needs an explicit closing tag;
             // it's Boolean false by default
           },
         ];
       }, */

      // optional; webpack config Object for
      // the custom service worker ONLY (/src-pwa/custom-service-worker.[js|ts])
      // if using workbox in InjectManifest mode
      extendWebpackCustomSW(cfg) {
        // directly change props of cfg;
        // no need to return anything
      },

      // optional; EQUIVALENT to extendWebpackCustomSW() but uses webpack-chain;
      // for the custom service worker ONLY (/src-pwa/custom-service-worker.[js|ts])
      // if using workbox in InjectManifest mode
      chainWebpackCustomSW(chain) {
        // chain is a webpack-chain instance
        // of the Webpack configuration
        // example:
        // chain.plugin('eslint-webpack-plugin')
        //   .use(ESLintPlugin, [{ extensions: [ 'js' ] }])
      },
    },

    // Full list of options: https://v2.quasar.dev/quasar-cli-vite/developing-cordova-apps/configuring-cordova
    cordova: {
      // noIosLegacyBuildFlag: true, // uncomment only if you know what you are doing
    },

    // Full list of options: https://v2.quasar.dev/quasar-cli-vite/developing-capacitor-apps/configuring-capacitor
    capacitor: {
      hideSplashscreen: true,
    },

    // Full list of options: https://v2.quasar.dev/quasar-cli-vite/developing-electron-apps/configuring-electron
    electron: {
      // extendElectronMainConf (esbuildConf)
      // extendElectronPreloadConf (esbuildConf)

      // specify the debugging port to use for the Electron app when running in development mode
      inspectPort: 5858,

      bundler: "packager", // 'packager' or 'builder'

      packager: {
        // https://github.com/electron-userland/electron-packager/blob/master/docs/api.md#options
        // OS X / Mac App Store
        // appBundleId: '',
        // appCategoryType: '',
        // osxSign: '',
        // protocol: 'myapp://path',
        // Windows only
        // win32metadata: { ... }
      },

      builder: {
        // https://www.electron.build/configuration/configuration

        appId: "quasar-project",
      },
    },

    // Full list of options: https://v2.quasar.dev/quasar-cli-vite/developing-browser-extensions/configuring-bex
    bex: {
      contentScripts: ["my-content-script"],

      // extendBexScriptsConf (esbuildConf) {}
      // extendBexManifestJson (json) {}
    },
  };
});

```

## default

```js
/* eslint-env node */

/*
 * This file runs in a Node context (it's NOT transpiled by Babel), so use only
 * the ES6 features that are supported by your Node version. https://node.green/
 */

// Configuration for your app
// https://v2.quasar.dev/quasar-cli-vite/quasar-config-js

const { configure } = require("quasar/wrappers");
const path = require("path");

module.exports = configure(function (/* ctx */) {
  return {
    // https://v2.quasar.dev/quasar-cli-vite/prefetch-feature
    // preFetch: true,

    // app boot file (/src/boot)
    // --> boot files are part of "main.js"
    // https://v2.quasar.dev/quasar-cli-vite/boot-files
    boot: ["i18n", "axios"],

    // https://v2.quasar.dev/quasar-cli-vite/quasar-config-js#css
    css: ["app.scss"],

    // https://github.com/quasarframework/quasar/tree/dev/extras
    extras: [
      // 'ionicons-v4',
      // 'mdi-v7',
      // 'fontawesome-v6',
      // 'eva-icons',
      // 'themify',
      // 'line-awesome',
      // 'roboto-font-latin-ext', // this or either 'roboto-font', NEVER both!

      "roboto-font", // optional, you are not bound to it
      "material-icons", // optional, you are not bound to it
    ],

    // Full list of options: https://v2.quasar.dev/quasar-cli-vite/quasar-config-js#build
    build: {
      target: {
        browser: ["es2019", "edge88", "firefox78", "chrome87", "safari13.1"],
        node: "node20",
      },

      vueRouterMode: "hash", // available values: 'hash', 'history'
      // vueRouterBase,
      // vueDevtools,
      // vueOptionsAPI: false,

      // rebuildCache: true, // rebuilds Vite/linter/etc cache on startup

      // publicPath: '/',
      // analyze: true,
      // env: {},
      // rawDefine: {}
      // ignorePublicFolder: true,
      // minify: false,
      // polyfillModulePreload: true,
      // distDir

      // extendViteConf (viteConf) {},
      // viteVuePluginOptions: {},

      vitePlugins: [
        [
          "@intlify/vite-plugin-vue-i18n",
          {
            // if you want to use Vue I18n Legacy API, you need to set `compositionOnly: false`
            // compositionOnly: false,

            // if you want to use named tokens in your Vue I18n messages, such as 'Hello {name}',
            // you need to set `runtimeOnly: false`
            // runtimeOnly: false,

            // you need to set i18n resource including paths !
            include: path.resolve(__dirname, "./src/i18n/**"),
          },
        ],
        [
          "vite-plugin-checker",
          {
            eslint: {
              lintCommand: 'eslint "./**/*.{js,mjs,cjs,vue}"',
            },
          },
          { server: false },
        ],
      ],
    },

    // Full list of options: https://v2.quasar.dev/quasar-cli-vite/quasar-config-js#devServer
    devServer: {
      // https: true
      open: true, // opens browser window automatically
    },

    // https://v2.quasar.dev/quasar-cli-vite/quasar-config-js#framework
    framework: {
      config: {},

      // iconSet: 'material-icons', // Quasar icon set
      // lang: 'en-US', // Quasar language pack

      // For special cases outside of where the auto-import strategy can have an impact
      // (like functional components as one of the examples),
      // you can manually specify Quasar components/directives to be available everywhere:
      //
      // components: [],
      // directives: [],

      // Quasar plugins
      plugins: [],
    },

    // animations: 'all', // --- includes all animations
    // https://v2.quasar.dev/options/animations
    animations: [],

    // https://v2.quasar.dev/quasar-cli-vite/quasar-config-js#property-sourcefiles
    // sourceFiles: {
    //   rootComponent: 'src/App.vue',
    //   router: 'src/router/index',
    //   store: 'src/store/index',
    //   registerServiceWorker: 'src-pwa/register-service-worker',
    //   serviceWorker: 'src-pwa/custom-service-worker',
    //   pwaManifestFile: 'src-pwa/manifest.json',
    //   electronMain: 'src-electron/electron-main',
    //   electronPreload: 'src-electron/electron-preload'
    // },

    // https://v2.quasar.dev/quasar-cli-vite/developing-ssr/configuring-ssr
    ssr: {
      // ssrPwaHtmlFilename: 'offline.html', // do NOT use index.html as name!
      // will mess up SSR

      // extendSSRWebserverConf (esbuildConf) {},
      // extendPackageJson (json) {},

      pwa: false,

      // manualStoreHydration: true,
      // manualPostHydrationTrigger: true,

      prodPort: 3000, // The default port that the production server should use
      // (gets superseded if process.env.PORT is specified at runtime)

      middlewares: [
        "render", // keep this as last one
      ],
    },

    // https://v2.quasar.dev/quasar-cli-vite/developing-pwa/configuring-pwa
    pwa: {
      workboxMode: "generateSW", // or 'injectManifest'
      injectPwaMetaTags: true,
      swFilename: "sw.js",
      manifestFilename: "manifest.json",
      useCredentialsForManifestTag: false,
      // useFilenameHashes: true,
      // extendGenerateSWOptions (cfg) {}
      // extendInjectManifestOptions (cfg) {},
      // extendManifestJson (json) {}
      // extendPWACustomSWConf (esbuildConf) {}
    },

    // Full list of options: https://v2.quasar.dev/quasar-cli-vite/developing-cordova-apps/configuring-cordova
    cordova: {
      // noIosLegacyBuildFlag: true, // uncomment only if you know what you are doing
    },

    // Full list of options: https://v2.quasar.dev/quasar-cli-vite/developing-capacitor-apps/configuring-capacitor
    capacitor: {
      hideSplashscreen: true,
    },

    // Full list of options: https://v2.quasar.dev/quasar-cli-vite/developing-electron-apps/configuring-electron
    electron: {
      // extendElectronMainConf (esbuildConf)
      // extendElectronPreloadConf (esbuildConf)

      // specify the debugging port to use for the Electron app when running in development mode
      inspectPort: 5858,

      bundler: "packager", // 'packager' or 'builder'

      packager: {
        // https://github.com/electron-userland/electron-packager/blob/master/docs/api.md#options
        // OS X / Mac App Store
        // appBundleId: '',
        // appCategoryType: '',
        // osxSign: '',
        // protocol: 'myapp://path',
        // Windows only
        // win32metadata: { ... }
      },

      builder: {
        // https://www.electron.build/configuration/configuration

        appId: "quasar-project",
      },
    },

    // Full list of options: https://v2.quasar.dev/quasar-cli-vite/developing-browser-extensions/configuring-bex
    bex: {
      contentScripts: ["my-content-script"],

      // extendBexScriptsConf (esbuildConf) {}
      // extendBexManifestJson (json) {}
    },
  };
});
```
