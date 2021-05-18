module.exports = {
    devServer: {
        proxy: {
            '/api': {
                target: 'http://127.0.0.1:9090',
                ws: true,
                changeOrigin: true
            }
        }
    },
    runtimeCompiler: true
}
