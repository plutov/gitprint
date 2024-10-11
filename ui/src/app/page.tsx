"use client";

import Form from "../components/Form";

export default function Home() {
  return (
    <div>
      <div className="pb-10 pt-10">
        <div className="mx-auto max-w-2xl px-4">
          <div className="flex flex-col gap-2 rounded-md border p-8">
            <h1 className="text-lg font-semibold">
              {">"} Print your favorite Git repositories as PDF
            </h1>
            <p className="py-2">
              Looking for a fun way to explore your favorite GitHub
              repositories? Tired of staring at the screen for hours on end? Or
              maybe want to print out a hard copy as a keepsake?
            </p>
            <p className="py-2">
              Simply sign in with your GitHub account and start printing public
              or private repositories in a beautiful, easy-to-read format.
            </p>
            <p className="py-2">
              It is currently in beta, so please be patient with us as we work.
              Feel free to request features or report bugs.
            </p>
            <p className="py-2 font-bold">
              Made by <a href="https://x.com/pliutau">@pliutau</a> with ❤️
            </p>
          </div>
        </div>
      </div>
      <div className="pb-10">{Form()}</div>
      <div className="pb-10">
        <div className="mx-auto max-w-2xl px-4">
          <div className="flex flex-col text-center pb-4">
            <h2 className="text-lg font-semibold">Repositories of the day</h2>
            <span className="text-sm text-gray-400">
              Click to download the PDF
            </span>
          </div>
          <div className="mb-4 grid grid-cols-2 gap-2">
            <a
              className="cursor-pointer rounded-md border p-4"
              href={`${process.env.NEXT_PUBLIC_API_ADDR}/files?export_id=dabee9e7b0dab2f5bd2a1b4f5c79929054d0fbc50126afbcb6b5af77d5fd7203&ext=pdf`}
            >
              <div className="text-sm font-semibold">docker/scout-cli</div>
              <div className="text-sm text-white">v1.14.0</div>
              <div className="text-sm text-zinc-600">2.4MB</div>
            </a>
            <a
              className="cursor-pointer rounded-md border p-4"
              href={`${process.env.NEXT_PUBLIC_API_ADDR}/files?export_id=153157a6e609ba9b6e61ebd9f86780212c293581c3aead1b199f3229d8165757&ext=pdf`}
            >
              <div className="text-sm font-semibold">astral-sh/ruff-lsp</div>
              <div className="text-sm text-white">v0.0.57</div>
              <div className="text-sm text-zinc-600">2.3MB</div>
            </a>
            <a
              className="cursor-pointer rounded-md border p-4"
              href={`${process.env.NEXT_PUBLIC_API_ADDR}/files?export_id=6d0bdc0f9d43f1167e6bf1d4e8c68eb35b0f8f88d8cfce9f99028f57c0f24c8b&ext=pdf`}
            >
              <div className="text-sm font-semibold">binarly-io/efiXplorer</div>
              <div className="text-sm text-white">v6.0</div>
              <div className="text-sm text-zinc-600">1MB</div>
            </a>
            <a
              className="cursor-pointer rounded-md border p-4"
              href={`${process.env.NEXT_PUBLIC_API_ADDR}/files?export_id=5b8926f28a32abc09d21f5695574b99e3cd6831925c3626a12f66c19ee69babd&ext=pdf`}
            >
              <div className="text-sm font-semibold">redis/redis-vl-python</div>
              <div className="text-sm text-white">0.3.4</div>
              <div className="text-sm text-zinc-600">2.5MB</div>
            </a>
          </div>
        </div>
      </div>
    </div>
  );
}
