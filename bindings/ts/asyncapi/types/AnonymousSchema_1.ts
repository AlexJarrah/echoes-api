import AnonymousSchema_3 from './AnonymousSchema_3';
import AnonymousSchema_6 from './AnonymousSchema_6';
class AnonymousSchema_1 {
  private _id: string;
  private _reservedType: AnonymousSchema_3;
  private _data?: AnonymousSchema_6 | null;
  private _additionalProperties?: Map<string, any>;

  constructor(input: {
    id: string,
    reservedType: AnonymousSchema_3,
    data?: AnonymousSchema_6 | null,
    additionalProperties?: Map<string, any>,
  }) {
    this._id = input.id;
    this._reservedType = input.reservedType;
    this._data = input.data;
    this._additionalProperties = input.additionalProperties;
  }

  get id(): string { return this._id; }
  set id(id: string) { this._id = id; }

  get reservedType(): AnonymousSchema_3 { return this._reservedType; }
  set reservedType(reservedType: AnonymousSchema_3) { this._reservedType = reservedType; }

  get data(): AnonymousSchema_6 | null | undefined { return this._data; }
  set data(data: AnonymousSchema_6 | null | undefined) { this._data = data; }

  get additionalProperties(): Map<string, any> | undefined { return this._additionalProperties; }
  set additionalProperties(additionalProperties: Map<string, any> | undefined) { this._additionalProperties = additionalProperties; }
}
export default AnonymousSchema_1;
