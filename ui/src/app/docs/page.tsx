"use client";

export default function Docs() {
  return (
    <div>
      <div className="pb-10 pt-10">
        <div className="mx-auto max-w-2xl px-4">
          <div className="flex flex-col gap-2 rounded-md border p-8">
            <h1 className="text-lg font-semibold">
              {">"} Which files are printed?
            </h1>
            <p>
              To keep the size of the PDF small, only the whitelisted extensions
              are include. That includes programming languages, markup
              languages, and some other common file types. Files bigger than
              100KB are not included in the PDF.
            </p>
          </div>
          <div className="flex flex-col gap-2 rounded-md border p-8 mt-4">
            <h1 className="text-lg font-semibold">{">"} Limitations</h1>
            <p>
              Because it is a free service, there are some limitations: you can
              generate only one PDF per hour.
            </p>
          </div>
          <div className="flex flex-col gap-2 rounded-md border p-8 mt-4">
            <h1 className="text-lg font-semibold">{">"} Known Issues</h1>
            <p>1. It may not work well on huge repositories.</p>
            <p>2. No syntax highlighting, but supports Markdown.</p>
          </div>
        </div>
      </div>
    </div>
  );
}
