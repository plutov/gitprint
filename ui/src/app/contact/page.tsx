"use client";

export default function Contact() {
  return (
    <div>
      <div className="pb-10 pt-10">
        <div className="mx-auto max-w-2xl px-4">
          <div className="flex flex-col gap-2 rounded-md border p-8">
            <h1 className="text-lg font-semibold">{">"} Contact me</h1>
            <p>
              Feel free to request features or report bugs. You can reach me at{" "}
              <span className="text-teal-500">
                a dot pliutau at gmail dot com
              </span>{" "}
              or on X - <a href="https://x.com/pliutau">@pliutau</a>.
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}
