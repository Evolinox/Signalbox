import { defineConfig } from 'vitepress'
import { version } from '../../package.json'
import { linkChangelog, linkRelease, socialDiscord, socialGit } from "./meta";

// https://vitepress.dev/reference/site-config
export default defineConfig({
  srcDir: "content",
  base: "/Signalbox",
  lang: "en-US",
  title: "Signalbox",
  description: "Modern way to control your Modelrailroad",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Docs', link: '/docs/overview/introduction' },
      { text: 'Showcase', link: '/showcase' },
      {
        text: `v${version}`,
        items: [
          { text: 'Changelog', link: linkChangelog },
          { text: 'Release', link: linkRelease }
        ]
      }
    ],

    sidebar: [
      {
        text: 'Overview',
        icon: 'lucide:rocket',
        items: [
          { text: 'Introduction', link: '/docs/overview/introduction' },
          { text: 'Install and Setup', link: '/docs/overview/install-setup' },
          { text: 'Changelog', link: '/docs/overview/changelog' }
        ]
      },
      {
        text: 'Control Box',
        icon: 'lucide:box',
        items: [
          { text: 'Introduction', link: '/docs/control-box/introduction' }
        ]
      },
      {
        text: 'Remote Box',
        icon: 'lucide:box',
        items: [
          { text: 'Introduction', link: '/docs/remote-box/introduction' },
          { text: 'API', link: '/docs/remote-box/api' }
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: socialGit },
      { icon: 'discord', link: socialDiscord }
    ],

    search: {
      provider: 'local',
    }
  }
})
