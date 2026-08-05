import AnonymousSchema_36 from './AnonymousSchema_36';
import AnonymousSchema_39 from './AnonymousSchema_39';
class AnonymousSchema_34 {
  private _id: string;
  private _reservedType: AnonymousSchema_36;
  private _data?: AnonymousSchema_39 | null;
  private _additionalProperties?: Map<string, any>;

  constructor(input: {
    id: string,
    reservedType: AnonymousSchema_36,
    data?: AnonymousSchema_39 | null,
    additionalProperties?: Map<string, any>,
  }) {
    this._id = input.id;
    this._reservedType = input.reservedType;
    this._data = input.data;
    this._additionalProperties = input.additionalProperties;
  }

  get id(): string { return this._id; }
  set id(id: string) { this._id = id; }

  get reservedType(): AnonymousSchema_36 { return this._reservedType; }
  set reservedType(reservedType: AnonymousSchema_36) { this._reservedType = reservedType; }

  get data(): AnonymousSchema_39 | null | undefined { return this._data; }
  set data(data: AnonymousSchema_39 | null | undefined) { this._data = data; }

  get additionalProperties(): Map<string, any> | undefined { return this._additionalProperties; }
  set additionalProperties(additionalProperties: Map<string, any> | undefined) { this._additionalProperties = additionalProperties; }
}
export default AnonymousSchema_34;
