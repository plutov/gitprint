export async function call(path: string, init?: RequestInit) {
  try {
    if (init) {
      init["cache"] = "no-store";
    }

    const host = process.env.NEXT_PUBLIC_API_ADDR;

    console.log("calling", `${host}${path}`);
    const res = await fetch(`${host}${path}`, init);
    const data = await res.json();

    if (!data) {
      return {
        status: 500,
        error: "invalid response",
        data: {},
      };
    }

    if (res.status < 200 || res.status >= 300) {
      console.error("unable to call the api", data.message ?? res.status);
      return {
        status: res.status,
        error: data.message ?? "unable to call the api",
        data: {},
      };
    }

    return {
      status: res.status,
      error: "",
      data: data.data,
    };
  } catch (e) {
    console.error("unable to call the api", e);
    return {
      status: 500,
      error: "unable to call the api",
      data: {},
    };
  }
}

export async function post(path: string, payload: object, jwt?: string) {
  const headers = {
    "Content-Type": "application/json",
    Authorization: "",
  };
  if (jwt) {
    headers["Authorization"] = `Bearer ${jwt}`;
  }

  return call(path, {
    method: "POST",
    body: JSON.stringify(payload),
    headers: headers,
  });
}

export async function get(path: string, jwt?: string) {
  const headers = {
    "Content-Type": "application/json",
    Authorization: "",
  };
  if (jwt) {
    headers["Authorization"] = `Bearer ${jwt}`;
  }

  return call(path, {
    method: "GET",
    headers: headers,
  });
}

export async function getJWT(code: string, state: string) {
  return await get(
    `/github/auth/callback?code=${encodeURIComponent(code)}&state=${encodeURIComponent(state)}`,
  );
}

export async function download(
  jwt: string,
  repo: string,
  ref: string,
  exclude: string,
) {
  return await get(
    `/private/github/repo/download?repo=${encodeURIComponent(repo)}&ref=${encodeURIComponent(ref)}&exclude=${encodeURIComponent(exclude)}`,
    jwt,
  );
}

export async function generate(
  jwt: string,
  repo: string,
  ref: string,
  exportId: string,
) {
  return await get(
    `/private/github/repo/generate?repo=${encodeURIComponent(repo)}&ref=${encodeURIComponent(ref)}&export_id=${encodeURIComponent(exportId)}`,
    jwt,
  );
}
