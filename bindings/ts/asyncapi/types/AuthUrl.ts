
class AuthUrl {
  private _url: string;
  private _additionalProperties?: Map<string, any>;

  constructor(input: {
    url: string,
    additionalProperties?: Map<string, any>,
  }) {
    this._url = input.url;
    this._additionalProperties = input.additionalProperties;
  }

  get url(): string { return this._url; }
  set url(url: string) { this._url = url; }

  get additionalProperties(): Map<string, any> | undefined { return this._additionalProperties; }
  set additionalProperties(additionalProperties: Map<string, any> | undefined) { this._additionalProperties = additionalProperties; }
}
export default AuthUrl;
