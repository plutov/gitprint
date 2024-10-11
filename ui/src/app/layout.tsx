import type { Metadata } from "next";
import { GoogleAnalytics } from "@next/third-parties/google";
import "./globals.css";

export const metadata: Metadata = {
  metadataBase: new URL("https://gitprint.me"),
  title: "Git Print",
  description: "Print your favorite Git repositories as PDF",
  openGraph: {
    type: "website",
    locale: "en_US",
    title: "Git Print",
    description: "Print your favorite Git repositories as PDF",
    siteName: "gitprint.me",
    images: [
      {
        url: "/logo.png",
      },
    ],
  },
  icons: {
    icon: "/favicon.ico",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="dark">
      <body className="flex flex-col h-screen justify-between">
        <header className="flex h-16 w-full items-center justify-center border-b">
          <span className="mx-2">
            <a href="/">home</a>
          </span>
          <span>•</span>
          <span className="mx-2">
            <a href="/docs">docs</a>
          </span>
          <span>•</span>
          <span className="mx-2">
            <a href="/contact">contact</a>
          </span>
        </header>
        <main className="flex flex-col mb-auto">{children}</main>
        <footer className="flex w-full items-center justify-center border-t">
          <div className="flex justify-center p-4 text-sm">
            <span className="mx-2">v0.1.1</span>
            <span>•</span>
            <span className="mx-2">
              Made by{" "}
              <a rel="nofollow" href="https://x.com/pliutau">
                @pliutau
              </a>
            </span>
            <span>•</span>
            <span className="mx-2">
              <a href="/docs">How it works</a>
            </span>
          </div>
        </footer>
      </body>
      <GoogleAnalytics gaId="G-59TJWLG7VY" />
    </html>
  );
}
