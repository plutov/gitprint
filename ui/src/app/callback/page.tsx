"use client";

import { getJWT } from "../../lib/api";
import { useSearchParams } from "next/navigation";
import { useEffect } from "react";

export default function Callback() {
  const searchParams = useSearchParams();

  const code = searchParams.get("code") as string;
  const state = searchParams.get("state") as string;

  // redirect to home if auth is not successful in 5s
  useEffect(() => {
    const timeout = setTimeout(() => {
      // window.location.href = "/";
    }, 5000);

    return () => {
      clearTimeout(timeout);
    };
  }, []);

  useEffect(() => {
    const fetchJWT = async () => {
      const res = await getJWT(code, state);
      if (!res.error && res.data.jwt_token) {
        localStorage.setItem(`auth_jwt`, res.data.jwt_token);
        window.location.href = "/";
      }
    };

    if (code && state) {
      fetchJWT();
    }
  }, [code, state]);

  return (
    <div>
      <div className="pt-20">
        <div className="mx-auto max-w-xl px-4">
          <div className="flex flex-col rounded-md border p-8 text-center">
            <p className="py-2">Authenticating...</p>
          </div>
        </div>
        <div className="h-px w-full"></div>
      </div>
    </div>
  );
}
