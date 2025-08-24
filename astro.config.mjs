import { defineConfig, passthroughImageService } from "astro/config";
import starlight from "@astrojs/starlight";
import starlightGiscus from "starlight-giscus";

// https://astro.build/config
export default defineConfig({
  site: "https://the.graphics.gd/",
  redirects: {
    "/": "/guide",
  },
  image: {
    service: passthroughImageService(),
  },
  integrations: [
    starlight({
      title: "the.graphics.gd/guide",

      social: [
        {
          icon: "github",
          label: "Github",
          href: "https://github.com/quaadgras/graphics.gd",
        },
      ],
      plugins: [
        starlightGiscus({
          repo: "quaadgras/graphics.gd",
          repoId: "R_kgDOH2vsEw",
          category: "Documentation",
          categoryId: "DIC_kwDOH2vsE84Cuhec",
        }),
      ],
      customCss: ["./src/styles/iframe.css"],
      sidebar: [
        {
          label: "Guide",
          autogenerate: { directory: "guide" },
        },
        {
          label: "Reference",
          items: [
            {
              label: "startup",
              link: "https://pkg.go.dev/graphics.gd/startup",
            },
            {
              label: "classdb",
              link: "https://pkg.go.dev/graphics.gd/classdb",
            },
            {
              label: "shaders",
              link: "https://pkg.go.dev/graphics.gd/shaders",
            },
            {
              label: "variant",
              link: "https://pkg.go.dev/graphics.gd/variant",
            },
          ],
        },
        {
          label: "License",
          autogenerate: { directory: "license" },
        },
      ],
    }),
  ],
});
