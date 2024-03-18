import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import postcss from "postcss";

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    cssCodeSplit: false,
    rollupOptions: {
      external: ['vue'],
      output: {
        format: 'umd',
        globals: {
          vue: 'Vue'
        },
        entryFileNames: `[name].js`,
        chunkFileNames: `[name].js`,
        assetFileNames: `[name].[ext]`
      }
    }
  },
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
})
