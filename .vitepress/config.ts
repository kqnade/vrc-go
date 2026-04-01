import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'vrc-go',
  description: 'VRChat API unofficial Go client library',
  base: '/vrcgo/',

  locales: {
    root: {
      label: 'English',
      lang: 'en',
      themeConfig: {
        nav: [
          { text: 'Guide', link: '/en/getting-started' },
          { text: 'API Reference', link: '/en/api/authentication' },
          {
            text: 'GitHub',
            link: 'https://github.com/kqnade/vrcgo',
          },
        ],
        sidebar: {
          '/en/': [
            {
              text: 'Introduction',
              items: [
                { text: 'Getting Started', link: '/en/getting-started' },
              ],
            },
            {
              text: 'API Reference',
              items: [
                { text: 'Authentication', link: '/en/api/authentication' },
                { text: 'Users', link: '/en/api/users' },
                { text: 'Friends', link: '/en/api/friends' },
                { text: 'Avatars', link: '/en/api/avatars' },
                { text: 'Worlds', link: '/en/api/worlds' },
                { text: 'Instances', link: '/en/api/instances' },
                { text: 'Notifications', link: '/en/api/notifications' },
                { text: 'Favorites', link: '/en/api/favorites' },
                { text: 'Groups', link: '/en/api/groups' },
                { text: 'Files', link: '/en/api/files' },
                { text: 'Player Moderation', link: '/en/api/player-moderation' },
                { text: 'System', link: '/en/api/system' },
                { text: 'WebSocket', link: '/en/api/websocket' },
              ],
            },
          ],
        },
      },
    },
    ja: {
      label: '日本語',
      lang: 'ja',
      themeConfig: {
        nav: [
          { text: 'ガイド', link: '/ja/getting-started' },
          { text: 'APIリファレンス', link: '/ja/api/authentication' },
          {
            text: 'GitHub',
            link: 'https://github.com/kqnade/vrcgo',
          },
        ],
        sidebar: {
          '/ja/': [
            {
              text: 'はじめに',
              items: [
                { text: 'Getting Started', link: '/ja/getting-started' },
              ],
            },
            {
              text: 'APIリファレンス',
              items: [
                { text: '認証', link: '/ja/api/authentication' },
                { text: 'ユーザー', link: '/ja/api/users' },
                { text: 'フレンド', link: '/ja/api/friends' },
                { text: 'アバター', link: '/ja/api/avatars' },
                { text: 'ワールド', link: '/ja/api/worlds' },
                { text: 'インスタンス', link: '/ja/api/instances' },
                { text: '通知', link: '/ja/api/notifications' },
                { text: 'お気に入り', link: '/ja/api/favorites' },
                { text: 'グループ', link: '/ja/api/groups' },
                { text: 'ファイル', link: '/ja/api/files' },
                { text: 'プレイヤーモデレーション', link: '/ja/api/player-moderation' },
                { text: 'システム', link: '/ja/api/system' },
                { text: 'WebSocket', link: '/ja/api/websocket' },
              ],
            },
          ],
        },
      },
    },
  },

  themeConfig: {
    logo: '/logo.svg',
    socialLinks: [
      { icon: 'github', link: 'https://github.com/kqnade/vrcgo' },
    ],
    footer: {
      message: 'Released under the Apache-2.0 License.',
      copyright: 'Copyright © 2024 kqnade',
    },
    search: {
      provider: 'local',
    },
  },
})
