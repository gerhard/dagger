import type { Config } from "@docusaurus/types";
import type * as Preset from "@docusaurus/preset-classic";
import { themes as prismThemes } from "prism-react-renderer";
import remarkCodeImport from "remark-code-import";
import remarkTemplate from "./plugins/remark-template";
import llmsTxtPlugin from "./plugins/llms-txt-plugin";

import { daggerVersion } from "./current_docs/partials/version";

const url = "https://docs.dagger.io";

const config: Config = {
  title: "Dagger",
  tagline:
    "Open-source runtime for composable workflows, powering AI agents and CI/CD with modular, repeatable, and observable pipelines.",
  favicon: "img/favicon.svg",

  // Set the production url of your site here
  url: url,
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: "/",

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  //organizationName: 'facebook', // Usually your GitHub org/user name.
  //projectName: 'docusaurus', // Usually your repo name.

  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "throw",

  // Even if you don't use internationalization, you can use this field to set
  // useful metadata like html lang. For example, if your site is Chinese, you
  // may want to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: "en",
    locales: ["en"],
  },
  markdown: {
    mermaid: true,
  },
  scripts: [
    {
      src: "/js/commonroom.js",
      async: true,
    },
  ],
  presets: [
    [
      "classic",
      {
        docs: {
          breadcrumbs: false,
          path: "./current_docs",
          routeBasePath: "/",
          sidebarPath: "./sidebars.ts",
          editUrl: "https://github.com/dagger/dagger/edit/main/docs",
          remarkPlugins: [
            [remarkCodeImport, { allowImportingFromOutside: true }],
            [remarkTemplate, { version: daggerVersion }],
          ],
        },
        blog: false,
        theme: {
          customCss: require.resolve("./src/css/custom.scss"),
        },
      } satisfies Preset.Options,
    ],
  ],
  plugins: [
    "docusaurus-plugin-sass",
    "docusaurus-plugin-image-zoom",
    // Thanks to @jharrell and Prisma team. Apache-2.0 content
    llmsTxtPlugin,
    [
      "posthog-docusaurus",
      {
        apiKey: "phc_rykA1oJnBnxTwavpgJKr4RAVXEgCkpyPVi21vQ7906d",
        appUrl: "https://us.i.posthog.com", // Changed to standard PostHog URL
        enableInDevelopment: true, // Enable tracking in development
      },
    ],
    [
      "docusaurus-plugin-typedoc",
      {
        id: "current-generation",
        plugin: ["typedoc-plugin-markdown", "typedoc-plugin-frontmatter"],
        entryPoints: [
          "../sdk/typescript/src/connect.ts",
          "../sdk/typescript/src/api/client.gen.ts",
          "../sdk/typescript/src/common/errors/index.ts",
        ],
        tsconfig: "../sdk/typescript/tsconfig.json",
        out: "current_docs/reference/typescript/",
        excludeProtected: true,
        exclude: "../sdk/typescript/node_modules/**",
        skipErrorChecking: true,
        disableSources: true,
        sanitizeComments: true,
        frontmatterGlobals: {
          displayed_sidebar: "current",
          sidebar_label: "TypeScript SDK Reference",
          title: "TypeScript SDK Reference",
        },
        textContentMappings: {
          "title.indexPage": "TypeScript SDK Reference",
          "footer.text": "",
        },
        requiredToBeDocumented: ["Class"],
      },
    ],
  ],
  themes: ["@docusaurus/theme-mermaid"],
  themeConfig: {
    sidebarCollapsed: false,
    metadata: [
      {
        name: "description",
        content:
          "Dagger is an open-source runtime for composable workflows, powering AI agents and CI/CD with modular, repeatable, and observable pipelines.",
      },
      {
        name: "image",
        property: "og:image",
        content: `${url}/img/current_docs/index/daggernaut-carpenter-robots-share.jpg`,
      },
      {
        name: "author",
        content: "Dagger",
      },
      {
        property: "twitter:image",
        content: `${url}/img/current_docs/index/daggernaut-carpenter-robots-share.jpg`,
      },
    ],
    prism: {
      additionalLanguages: [
        "php",
        "rust",
        "elixir",
        "bash",
        "toml",
        "powershell",
        "java",
      ],
      theme: prismThemes.dracula,
    },
    navbar: {
      logo: {
        alt: "Dagger Logo",
        src: "img/dagger-logo-white.svg",
        height: "50px",
        href: "https://dagger.io/",
      },
      items: [
        {
          position: "left",
          to: "https://daggerverse.dev/",
          label: "Daggerverse",
          className: "navbar-blog-link",
          target: "_self",
        },
        {
          position: "left",
          to: "/",
          label: "Docs",
          className: "navbar-blog-link",
        },
        {
          type: "search",
          position: "right",
          className: "header-searchbar",
        },
      ],
    },
    algolia: {
      apiKey: "bffda1490c07dcce81a26a144115cc02",
      indexName: "dagger",
      appId: "XEIYPBWGOI",
    },
    colorMode: {
      defaultMode: "light",
    },
    zoom: {
      selector: ".markdown img:not(.not-zoom)",
      background: {
        light: "rgb(255, 255, 255)",
        dark: "rgb(50, 50, 50)",
      },
      // medium-zoom configuration options
      // Refer to https://github.com/francoischalifour/medium-zoom#options
      config: {},
    },
    footer: {
      links: [
        {
          title: "Resources",
          items: [
            {
              label: "Case Studies",
              to: "https://dagger.io/case-studies",
            },
            {
              label: "Videos",
              to: "https://dagger.io/videos",
            },
            {
              label: "Adopting Dagger",
              to: "https://dagger.io/adopting-dagger",
            },
            {
              label: "Daggerized Projects",
              to: "https://dagger.io/daggerized-projects",
            },
            {
              label: "Docs",
              to: "https://docs.dagger.io/",
            },
            {
              label: "Blog",
              to: "https://dagger.io/blog",
            },
            {
              label: "Community Content",
              to: "https://dagger.io/community-content",
            },
          ],
        },
        {
          title: "Community",
          items: [
            {
              label: "Events",
              to: "https://dagger.io/events",
            },
            {
              label: "Get Involved",
              to: "https://dagger.io/community",
            },
            {
              label: "Dagger Love",
              to: "https://dagger.io/dagger-love",
            },
            {
              label: "Dagger Commanders",
              to: "https://dagger.io/commanders",
            },
          ],
        },
        {
          title: "Product",
          items: [
            {
              label: "Dagger Engine",
              to: "https://dagger.io/dagger-engine",
            },
            {
              label: "Dagger Cloud",
              to: "https://dagger.io/cloud",
            },
            {
              label: "Daggerverse",
              to: "https://daggerverse.dev",
            },
            {
              label: "Integrations",
              to: "https://dagger.io/integrations",
            },
            {
              label: "Pricing",
              to: "https://dagger.io/pricing",
            },
          ],
        },
        {
          title: "Company",
          items: [
            {
              label: "Partners",
              to: "https://dagger.io/partners",
            },
            {
              label: "Careers",
              to: "https://boards.greenhouse.io/dagger",
            },
            {
              label: "Brand",
              to: "https://dagger.io/brand",
            },
            {
              label: "Terms of Service",
              to: "https://dagger.io/terms-of-service",
            },
            {
              label: "Privacy Policy",
              to: "https://dagger.io/privacy-policy",
            },
            {
              label: "Trademark Guidelines",
              to: "https://dagger.io/trademark-guidelines",
            },
            {
              label: "Dagger Trust Center",
              to: "https://trust.dagger.io",
            },
          ],
        },
      ],
      copyright: `
        <hr />
        <div class="flex justify-between">
          <small>© Dagger 2022-2024</small>
          <div class="flex gap-8">
              <a target="_blank" class="footer-discord-link" href="https://discord.gg/dagger-io">
              </a>
              <a target="_blank" class="footer-x-link" href="https://twitter.com/dagger_io">
              </a>
          </div>
        </div>
      `,
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
