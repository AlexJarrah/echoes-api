import AnonymousSchema_44 from './AnonymousSchema_44';
import AnonymousSchema_48 from './AnonymousSchema_48';
import AnonymousSchema_63 from './AnonymousSchema_63';
import AnonymousSchema_65 from './AnonymousSchema_65';
class AnonymousSchema_42 {
  private _id: string;
  private _reservedType: AnonymousSchema_44;
  private _data: Map<string, AnonymousSchema_48[]> | AnonymousSchema_63 | AnonymousSchema_65;
  private _additionalProperties?: Map<string, any>;

  constructor(input: {
    id: string,
    reservedType: AnonymousSchema_44,
    data: Map<string, AnonymousSchema_48[]> | AnonymousSchema_63 | AnonymousSchema_65,
    additionalProperties?: Map<string, any>,
  }) {
    this._id = input.id;
    this._reservedType = input.reservedType;
    this._data = input.data;
    this._additionalProperties = input.additionalProperties;
  }

  get id(): string { return this._id; }
  set id(id: string) { this._id = id; }

  get reservedType(): AnonymousSchema_44 { return this._reservedType; }
  set reservedType(reservedType: AnonymousSchema_44) { this._reservedType = reservedType; }

  get data(): Map<string, AnonymousSchema_48[]> | AnonymousSchema_63 | AnonymousSchema_65 { return this._data; }
  set data(data: Map<string, AnonymousSchema_48[]> | AnonymousSchema_63 | AnonymousSchema_65) { this._data = data; }

  get additionalProperties(): Map<string, any> | undefined { return this._additionalProperties; }
  set additionalProperties(additionalProperties: Map<string, any> | undefined) { this._additionalProperties = additionalProperties; }
}
export default AnonymousSchema_42;
