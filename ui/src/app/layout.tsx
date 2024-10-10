import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  metadataBase: new URL("https://gitprint.me"),
  title: "gitprint.me",
  description: "Print your favorite Git repositories as PDF",
  openGraph: {
    type: "website",
    locale: "en_US",
    title: "gitprint.me",
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
          <a rel="nofollow" href="/" className="font-bold p-8">
            gitprint.me - Print your favorite Git repositories as PDF.
          </a>
        </header>
        <main className="flex flex-col mb-auto">{children}</main>
        <footer className="flex w-full items-center justify-center border-t">
          <div className="flex justify-center p-4 text-sm">
            <span className="mx-2">v0.1.0</span>
            <span className="mx-2">•</span>
            <span className="mx-2">
              Made by{" "}
              <a rel="nofollow" href="https://x.com/pliutau">
                @pliutau
              </a>
            </span>
            <span className="mx-2">•</span>
            <a rel="nofollow" href="/docs">
              How it works
            </a>
          </div>
        </footer>
      </body>
    </html>
  );
}
