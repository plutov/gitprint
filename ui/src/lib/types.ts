export type User = {
  username: string;
};

export type JWTClaims = {
  user: User;
  exp: number;
};
