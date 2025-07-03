import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  optimizeDeps: {
    exclude: ['lucide-react'],
  },
  server: {
    proxy: {
      '/users': {
        target: 'http://localhost:8080', // Change this if your backend runs elsewhere
        changeOrigin: true,
      },
    },
  },
});
