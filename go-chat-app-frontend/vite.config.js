import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server:{
    proxy:{
      "/users":{
        target:"http://localhost:8080",
        changeOrigin:true,
      },
      "/login":{
        target:"http://localhost:8080",
        changeOrigin:true,
      },
      "/channels":{
        target:"http://localhost:8080",
        changeOrigin:true,
      },
      "/messages":{
        target:"http://localhost:8080",
        changeOrigin:true,
      },
    }
  }
})
