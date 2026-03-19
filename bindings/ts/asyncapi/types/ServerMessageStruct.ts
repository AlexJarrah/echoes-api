import ServerEventType from './ServerEventType';
import Activity from './Activity';
import AuthUrl from './AuthUrl';
import AuthToken from './AuthToken';
class ServerMessageStruct {
  private _id: string;
  private _reservedType: ServerEventType;
  private _data: Map<string, Activity[]> | AuthUrl | AuthToken;
  private _additionalProperties?: Map<string, any>;

  constructor(input: {
    id: string,
    reservedType: ServerEventType,
    data: Map<string, Activity[]> | AuthUrl | AuthToken,
    additionalProperties?: Map<string, any>,
  }) {
    this._id = input.id;
    this._reservedType = input.reservedType;
    this._data = input.data;
    this._additionalProperties = input.additionalProperties;
  }

  get id(): string { return this._id; }
  set id(id: string) { this._id = id; }

  get reservedType(): ServerEventType { return this._reservedType; }
  set reservedType(reservedType: ServerEventType) { this._reservedType = reservedType; }

  get data(): Map<string, Activity[]> | AuthUrl | AuthToken { return this._data; }
  set data(data: Map<string, Activity[]> | AuthUrl | AuthToken) { this._data = data; }

  get additionalProperties(): Map<string, any> | undefined { return this._additionalProperties; }
  set additionalProperties(additionalProperties: Map<string, any> | undefined) { this._additionalProperties = additionalProperties; }
}
export default ServerMessageStruct;
