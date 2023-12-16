import fs from 'fs';
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import { resolve } from 'path';

let exitHandlersBound = false;
const baseDir = '';
const buildDir = resolve(__dirname, 'public/build');
const hotFile = resolve(__dirname, 'public/hot');

const inertiaFiber = (rollupInput) => ({
    name: 'inertia-fiber',
    config: () => {
        return {
            base: baseDir,
            publicDir: false,
            refresh: true,
            build: {
                manifest: 'manifest.json',
                outDir: buildDir,
                rollupOptions: {
                    input: rollupInput,
                },
            },
            server: {
                cors: true,
                host: '127.0.1.1',
                port: 5173,
                strictPort: true,
                hmr: true,
            },
            resolve: {
                alias: {
                    '@': resolve(__dirname, 'resources/js'),
                },
            },
        }
    },
    configureServer(server) {
        server.httpServer?.once('listening', () => {
            const viteAddress = server.httpServer?.address();

            if (viteAddress) {
                const viteDevServerUrl = `http://${viteAddress.address}:${viteAddress.port}`;
                fs.writeFileSync(hotFile, viteDevServerUrl);

                setTimeout(() => {
                    server.config.logger.info('\n➜ Artefak');
                }, 100);
            }
        })

        if (! exitHandlersBound) {
            const clean = () => {
                if (fs.existsSync(hotFile)) {
                    fs.rmSync(hotFile);
                }
            }

            process.on('exit', clean);
            process.on('SIGINT', process.exit);
            process.on('SIGTERM', process.exit);
            process.on('SIGHUP', process.exit);

            exitHandlersBound = true;
        }

        return () => server.middlewares.use((next) => {
            next();
        })
    },
})

export default defineConfig({
    plugins: [
        inertiaFiber('resources/js/app.js'),
        vue({
            template: {
                transformAssetUrls: {
                    base: null,
                    includeAbsolute: false,
                },
            },
        }),
    ],
});
