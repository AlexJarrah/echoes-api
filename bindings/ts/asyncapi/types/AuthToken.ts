
class AuthToken {
  private _token: string;
  private _additionalProperties?: Map<string, any>;

  constructor(input: {
    token: string,
    additionalProperties?: Map<string, any>,
  }) {
    this._token = input.token;
    this._additionalProperties = input.additionalProperties;
  }

  get token(): string { return this._token; }
  set token(token: string) { this._token = token; }

  get additionalProperties(): Map<string, any> | undefined { return this._additionalProperties; }
  set additionalProperties(additionalProperties: Map<string, any> | undefined) { this._additionalProperties = additionalProperties; }
}
export default AuthToken;
