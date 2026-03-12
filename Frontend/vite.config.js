import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      '/server': {
        target: 'http://localhost:8000',
        changeOrigin: true,
        secure: false,
        rewrite: (path) => {
          console.log("[DEBUG] Proxying request:", path)
          return path.replace(/^\/server/, '')
        }
      },
    },
  },
})
