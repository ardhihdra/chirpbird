module.exports = function (api) {
    api.cache(true);
    const presets = ["module:metro-react-native-babel-preset"]
    const plugins = [
        [
            require.resolve('babel-plugin-module-resolver'),
            {
                "root": ["./src/"],
                "extensions": [
                    ".ios.ts",
                    ".android.ts",
                    ".ts",
                    ".ios.tsx",
                    ".android.tsx",
                    ".tsx",
                    ".jsx",
                    ".js",
                    ".json"
                ],
                "alias": {
                    "@core": "./src/core",
                    "@components": "./src/components",
                    "@assets": "./src/assets"
                }
            }
        ]
    ]
    return {
        presets,
        plugins
    }
}