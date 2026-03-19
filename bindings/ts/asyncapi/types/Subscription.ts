
class Subscription {
  private _users: string[];
  private _additionalProperties?: Map<string, any>;

  constructor(input: {
    users: string[],
    additionalProperties?: Map<string, any>,
  }) {
    this._users = input.users;
    this._additionalProperties = input.additionalProperties;
  }

  get users(): string[] { return this._users; }
  set users(users: string[]) { this._users = users; }

  get additionalProperties(): Map<string, any> | undefined { return this._additionalProperties; }
  set additionalProperties(additionalProperties: Map<string, any> | undefined) { this._additionalProperties = additionalProperties; }
}
export default Subscription;
