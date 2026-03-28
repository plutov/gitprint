export type User = {
  username: string;
};

export type JWTClaims = {
  user: User;
  exp: number;
};

export type RepoInfo = {
  name: string;
  size: string;
  version: string;
  download_url: string;
};
