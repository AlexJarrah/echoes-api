import Artist from './Artist';
class Activity {
  private _sessionId: string;
  private _trackId: string;
  private _trackName: string;
  private _artists: Artist[];
  private _albumName?: string;
  private _albumImageUrl?: string;
  private _active: boolean;
  private _position: number;
  private _startTimestamp: string;
  private _endTimestamp: string;
  private _updatedAt: string;
  private _additionalProperties?: Map<string, any>;

  constructor(input: {
    sessionId: string,
    trackId: string,
    trackName: string,
    artists: Artist[],
    albumName?: string,
    albumImageUrl?: string,
    active: boolean,
    position: number,
    startTimestamp: string,
    endTimestamp: string,
    updatedAt: string,
    additionalProperties?: Map<string, any>,
  }) {
    this._sessionId = input.sessionId;
    this._trackId = input.trackId;
    this._trackName = input.trackName;
    this._artists = input.artists;
    this._albumName = input.albumName;
    this._albumImageUrl = input.albumImageUrl;
    this._active = input.active;
    this._position = input.position;
    this._startTimestamp = input.startTimestamp;
    this._endTimestamp = input.endTimestamp;
    this._updatedAt = input.updatedAt;
    this._additionalProperties = input.additionalProperties;
  }

  get sessionId(): string { return this._sessionId; }
  set sessionId(sessionId: string) { this._sessionId = sessionId; }

  get trackId(): string { return this._trackId; }
  set trackId(trackId: string) { this._trackId = trackId; }

  get trackName(): string { return this._trackName; }
  set trackName(trackName: string) { this._trackName = trackName; }

  get artists(): Artist[] { return this._artists; }
  set artists(artists: Artist[]) { this._artists = artists; }

  get albumName(): string | undefined { return this._albumName; }
  set albumName(albumName: string | undefined) { this._albumName = albumName; }

  get albumImageUrl(): string | undefined { return this._albumImageUrl; }
  set albumImageUrl(albumImageUrl: string | undefined) { this._albumImageUrl = albumImageUrl; }

  get active(): boolean { return this._active; }
  set active(active: boolean) { this._active = active; }

  get position(): number { return this._position; }
  set position(position: number) { this._position = position; }

  get startTimestamp(): string { return this._startTimestamp; }
  set startTimestamp(startTimestamp: string) { this._startTimestamp = startTimestamp; }

  get endTimestamp(): string { return this._endTimestamp; }
  set endTimestamp(endTimestamp: string) { this._endTimestamp = endTimestamp; }

  get updatedAt(): string { return this._updatedAt; }
  set updatedAt(updatedAt: string) { this._updatedAt = updatedAt; }

  get additionalProperties(): Map<string, any> | undefined { return this._additionalProperties; }
  set additionalProperties(additionalProperties: Map<string, any> | undefined) { this._additionalProperties = additionalProperties; }
}
export default Activity;
