import { JWTClaims } from "./types";

export function parseJwt(token: string): JWTClaims | undefined {
  if (!token) {
    return undefined;
  }

  const base64Url = token.split(".")[1];
  const base64 = base64Url.replace("-", "+").replace("_", "/");
  const jwt = JSON.parse(window.atob(base64)) as JWTClaims;

  if (jwt.exp && jwt.exp * 1000 < Date.UTC(Date.now())) {
    return undefined;
  }

  return jwt;
}
