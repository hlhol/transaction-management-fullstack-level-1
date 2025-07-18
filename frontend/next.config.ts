import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  async rewrites() {
    return [
      {
        source: '/transactions/:path*',
        destination: 'http://localhost:8080/transactions/:path*',
      },
      {
        source: '/accounts/:path*',
        destination: 'http://localhost:8080/accounts/:path*',
      },
    ];
  },
};

export default nextConfig;
