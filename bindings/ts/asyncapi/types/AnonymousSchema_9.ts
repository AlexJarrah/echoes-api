import AnonymousSchema_11 from './AnonymousSchema_11';
import AnonymousSchema_15 from './AnonymousSchema_15';
import AnonymousSchema_30 from './AnonymousSchema_30';
import AnonymousSchema_32 from './AnonymousSchema_32';
class AnonymousSchema_9 {
  private _id: string;
  private _reservedType: AnonymousSchema_11;
  private _data: Map<string, AnonymousSchema_15[]> | AnonymousSchema_30 | AnonymousSchema_32;
  private _additionalProperties?: Map<string, any>;

  constructor(input: {
    id: string,
    reservedType: AnonymousSchema_11,
    data: Map<string, AnonymousSchema_15[]> | AnonymousSchema_30 | AnonymousSchema_32,
    additionalProperties?: Map<string, any>,
  }) {
    this._id = input.id;
    this._reservedType = input.reservedType;
    this._data = input.data;
    this._additionalProperties = input.additionalProperties;
  }

  get id(): string { return this._id; }
  set id(id: string) { this._id = id; }

  get reservedType(): AnonymousSchema_11 { return this._reservedType; }
  set reservedType(reservedType: AnonymousSchema_11) { this._reservedType = reservedType; }

  get data(): Map<string, AnonymousSchema_15[]> | AnonymousSchema_30 | AnonymousSchema_32 { return this._data; }
  set data(data: Map<string, AnonymousSchema_15[]> | AnonymousSchema_30 | AnonymousSchema_32) { this._data = data; }

  get additionalProperties(): Map<string, any> | undefined { return this._additionalProperties; }
  set additionalProperties(additionalProperties: Map<string, any> | undefined) { this._additionalProperties = additionalProperties; }
}
export default AnonymousSchema_9;
