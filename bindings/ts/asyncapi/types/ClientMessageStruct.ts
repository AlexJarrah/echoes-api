import ClientEventType from './ClientEventType';
import Subscription from './Subscription';
class ClientMessageStruct {
  private _id: string;
  private _reservedType: ClientEventType;
  private _data?: Subscription | null;
  private _additionalProperties?: Map<string, any>;

  constructor(input: {
    id: string,
    reservedType: ClientEventType,
    data?: Subscription | null,
    additionalProperties?: Map<string, any>,
  }) {
    this._id = input.id;
    this._reservedType = input.reservedType;
    this._data = input.data;
    this._additionalProperties = input.additionalProperties;
  }

  get id(): string { return this._id; }
  set id(id: string) { this._id = id; }

  get reservedType(): ClientEventType { return this._reservedType; }
  set reservedType(reservedType: ClientEventType) { this._reservedType = reservedType; }

  get data(): Subscription | null | undefined { return this._data; }
  set data(data: Subscription | null | undefined) { this._data = data; }

  get additionalProperties(): Map<string, any> | undefined { return this._additionalProperties; }
  set additionalProperties(additionalProperties: Map<string, any> | undefined) { this._additionalProperties = additionalProperties; }
}
export default ClientMessageStruct;
