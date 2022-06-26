module.exports = {
    chainWebpack: config => {
        config.resolve.alias
            .set("@", resolve("src"))
    },
    devServer: {
        proxy: {
            '/api': {
                target: 'http://47.103.81.142:8080',  // target host
                ws: true,  // proxy websockets 
                changeOrigin: true,  // needed for virtual hosted sites
                pathRewrite: {
                    '^/api': ''  // rewrite path
                }
            },
        }
    }  
}