/** @type {import('next').NextConfig} */
const nextConfig = {
  images: {
    remotePatterns: [
      {
        hostname: 'ucarecdn.com',
        protocol: 'https',
      },
    ],
  },
}

module.exports = nextConfig
