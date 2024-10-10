"use client";

import Form from "../components/Form";

export default function Home() {
  return (
    <div>
      <div className="pb-10 pt-10">
        <div className="mx-auto max-w-2xl px-4">
          <div className="flex flex-col gap-2 rounded-md border p-8">
            <h1 className="text-lg font-semibold">
              {">"} Welcome to gitprint.me
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
            <h1 className="text-lg font-semibold">Repositories of the day</h1>
            <span className="text-sm text-gray-400">
              Click to download the PDF
            </span>
          </div>
          <div className="mb-4 grid grid-cols-2 gap-2">
            <a
              className="cursor-pointer rounded-md border p-4"
              href={`${process.env.NEXT_PUBLIC_API_URL}/files?export_id=&ext=pdf`}
            >
              <div className="text-sm font-semibold">docker/scout-cli</div>
              <div className="text-sm text-white">v1.14.0</div>
              <div className="text-sm text-zinc-600">2.4MB</div>
            </a>
            <a
              className="cursor-pointer rounded-md border p-4"
              href={`${process.env.NEXT_PUBLIC_API_URL}/files?export_id=&ext=pdf`}
            >
              <div className="text-sm font-semibold">astral-sh/ruff-lsp</div>
              <div className="text-sm text-white">v0.0.57</div>
              <div className="text-sm text-zinc-600">2.3MB</div>
            </a>
            <a
              className="cursor-pointer rounded-md border p-4"
              href={`${process.env.NEXT_PUBLIC_API_URL}/files?export_id=&ext=pdf`}
            >
              <div className="text-sm font-semibold">binarly-io/efiXplorer</div>
              <div className="text-sm text-white">v6.0</div>
              <div className="text-sm text-zinc-600">1MB</div>
            </a>
            <a
              className="cursor-pointer rounded-md border p-4"
              href={`${process.env.NEXT_PUBLIC_API_URL}/files?export_id=&ext=pdf`}
            >
              <div className="text-sm font-semibold">redis/redis-vl-python</div>
              <div className="text-sm text-white">v0.3.4</div>
              <div className="text-sm text-zinc-600">2.5MB</div>
            </a>
          </div>
        </div>
      </div>
    </div>
  );
}
