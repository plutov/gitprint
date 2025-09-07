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
              href={`${process.env.NEXT_PUBLIC_API_ADDR}/files?export_id=fc6b4b70a97bb0d0d91773f443b472913a2ef34c623b670cc0e27326ceae70af&ext=pdf`}
            >
              <div className="text-sm font-semibold">plutov/practice-go</div>
              <div className="text-sm text-white">-</div>
              <div className="text-sm text-zinc-600">1.3MB</div>
            </a>
            <a
              className="cursor-pointer rounded-md border p-4"
              href={`${process.env.NEXT_PUBLIC_API_ADDR}/files?export_id=513d3bd286092b38054ae373e3f6051309956284b7ac4bc0873bd3450271c825&ext=pdf`}
            >
              <div className="text-sm font-semibold">google/mangle</div>
              <div className="text-sm text-white">v0.3.0</div>
              <div className="text-sm text-zinc-600">2.1MB</div>
            </a>
          </div>
        </div>
      </div>
    </div>
  );
}
