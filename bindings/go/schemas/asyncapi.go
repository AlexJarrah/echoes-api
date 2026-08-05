package schemas

const AsyncAPI = `asyncapi: 3.0.0
info:
  title: Echoes
  description: Real-time WebSocket API.
  version: 0.1.0
defaultContentType: application/json
servers:
  production:
    host: echoes.la
    protocol: ws
    description: Production WebSocket server.
channels:
  events:
    address: /api/events
    description: |
      Primary WebSocket channel for authenticated real-time events.

      Requires a valid ` + "`" + `auth_token` + "`" + ` session cookie.
    messages:
      subscribe:
        name: ClientMessage
        contentType: application/json
        payload:
          type: object
          required:
            - id
            - type
          properties:
            id:
              type: string
              format: uuid
              description: Unique correlation ID. Echoed in server responses.
            type:
              type: integer
              enum:
                - 0
                - 1
              description: |
                0 - AuthenticationRequest: initiate device authentication
                1 - ActivitySubscription: set user activity subscription
            data:
              oneOf:
                - type: 'null'
                - type: object
                  required:
                    - users
                  properties:
                    users:
                      type: array
                      items:
                        type: string
                        format: uuid
                      description: User UUIDs to subscribe to.
              description: Payload varies by type.
      activity:
        name: ServerMessage
        contentType: application/json
        payload:
          type: object
          required:
            - id
            - type
            - data
          properties:
            id:
              type: string
              format: uuid
              description: Correlation ID from the originating request.
            type:
              type: integer
              enum:
                - 0
                - 1
                - 2
              description: |
                0 - ActivityData: user activity snapshot
                1 - AuthenticationURL: one-time authentication URL
                2 - Authentication: session token
            data:
              oneOf:
                - type: object
                  additionalProperties:
                    type: array
                    items:
                      type: object
                      required:
                        - session_id
                        - track_id
                        - track_name
                        - artists
                        - active
                        - position
                        - start_timestamp
                        - end_timestamp
                        - updated_at
                      properties:
                        session_id:
                          type: string
                          description: Unique identifier for this listening session.
                        track_id:
                          type: string
                          format: uuid
                          description: Track ID.
                        track_name:
                          type: string
                          description: Track title.
                        artists:
                          type: array
                          items:
                            type: object
                            required:
                              - id
                              - name
                            properties:
                              id:
                                type: string
                                format: uuid
                                description: Artist ID.
                              name:
                                type: string
                                description: Artist name.
                          description: Credited track artists.
                        album_name:
                          type: string
                          nullable: true
                          description: Album name.
                        album_asset_url:
                          type: string
                          nullable: true
                          description: Album cover URL.
                        active:
                          type: boolean
                          description: >-
                            Whether the user is currently playing this track
                            (true) or has stopped/paused (false)..
                        position:
                          type: integer
                          format: int64
                          description: Current playback position in seconds.
                        start_timestamp:
                          type: string
                          format: date-time
                          description: When the track was played.
                        end_timestamp:
                          type: string
                          format: date-time
                          description: Projected end time based on track duration.
                        updated_at:
                          type: string
                          format: date-time
                          description: When this activity was last updated.
                  description: Map of user UUID to their activities.
                - type: object
                  required:
                    - url
                  properties:
                    url:
                      type: string
                      format: uri
                      description: One-time URL for the user to authorize the device.
                - type: object
                  required:
                    - token
                  properties:
                    token:
                      type: string
                      description: JWT authentication token.
              description: Payload varies by type.
  auth:
    address: /api/auth-events
    description: WebSocket channel for unauthenticated events.
    messages:
      request:
        name: ClientMessage
        contentType: application/json
        payload:
          type: object
          required:
            - id
            - type
          properties:
            id:
              type: string
              format: uuid
              description: Unique correlation ID. Echoed in server responses.
            type:
              type: integer
              enum:
                - 0
                - 1
              description: |
                0 - AuthenticationRequest: initiate device authentication
                1 - ActivitySubscription: set user activity subscription
            data:
              oneOf:
                - type: 'null'
                - type: object
                  required:
                    - users
                  properties:
                    users:
                      type: array
                      items:
                        type: string
                        format: uuid
                      description: User UUIDs to subscribe to.
              description: Payload varies by type.
      response:
        name: ServerMessage
        contentType: application/json
        payload:
          type: object
          required:
            - id
            - type
            - data
          properties:
            id:
              type: string
              format: uuid
              description: Correlation ID from the originating request.
            type:
              type: integer
              enum:
                - 0
                - 1
                - 2
              description: |
                0 - ActivityData: user activity snapshot
                1 - AuthenticationURL: one-time authentication URL
                2 - Authentication: session token
            data:
              oneOf:
                - type: object
                  additionalProperties:
                    type: array
                    items:
                      type: object
                      required:
                        - session_id
                        - track_id
                        - track_name
                        - artists
                        - active
                        - position
                        - start_timestamp
                        - end_timestamp
                        - updated_at
                      properties:
                        session_id:
                          type: string
                          description: Unique identifier for this listening session.
                        track_id:
                          type: string
                          format: uuid
                          description: Track ID.
                        track_name:
                          type: string
                          description: Track title.
                        artists:
                          type: array
                          items:
                            type: object
                            required:
                              - id
                              - name
                            properties:
                              id:
                                type: string
                                format: uuid
                                description: Artist ID.
                              name:
                                type: string
                                description: Artist name.
                          description: Credited track artists.
                        album_name:
                          type: string
                          nullable: true
                          description: Album name.
                        album_asset_url:
                          type: string
                          nullable: true
                          description: Album cover URL.
                        active:
                          type: boolean
                          description: >-
                            Whether the user is currently playing this track
                            (true) or has stopped/paused (false)..
                        position:
                          type: integer
                          format: int64
                          description: Current playback position in seconds.
                        start_timestamp:
                          type: string
                          format: date-time
                          description: When the track was played.
                        end_timestamp:
                          type: string
                          format: date-time
                          description: Projected end time based on track duration.
                        updated_at:
                          type: string
                          format: date-time
                          description: When this activity was last updated.
                  description: Map of user UUID to their activities.
                - type: object
                  required:
                    - url
                  properties:
                    url:
                      type: string
                      format: uri
                      description: One-time URL for the user to authorize the device.
                - type: object
                  required:
                    - token
                  properties:
                    token:
                      type: string
                      description: JWT authentication token.
              description: Payload varies by type.
operations:
  subscribeActivity:
    action: send
    channel:
      $ref: '#/channels/events'
    summary: Subscribe to user activity updates
    security:
      - type: httpApiKey
        in: cookie
        name: auth_token
        description: JWT token
    messages:
      - $ref: '#/channels/events/messages/subscribe'
  receiveActivity:
    action: receive
    channel:
      $ref: '#/channels/events'
    summary: Receive activity updates
    security:
      - type: httpApiKey
        in: cookie
        name: auth_token
        description: JWT token
    messages:
      - $ref: '#/channels/events/messages/activity'
  requestAuthentication:
    action: send
    channel:
      $ref: '#/channels/auth'
    summary: Request device authentication URL
    messages:
      - $ref: '#/channels/auth/messages/request'
  receiveAuthResponse:
    action: receive
    channel:
      $ref: '#/channels/auth'
    summary: Receive authentication URL and token
    messages:
      - $ref: '#/channels/auth/messages/response'
components:
  securitySchemes:
    CookieAuth:
      type: httpApiKey
      in: cookie
      name: auth_token
      description: JWT token
  messages:
    ClientMessage:
      name: ClientMessage
      contentType: application/json
      payload:
        type: object
        required:
          - id
          - type
        properties:
          id:
            type: string
            format: uuid
            description: Unique correlation ID. Echoed in server responses.
          type:
            type: integer
            enum:
              - 0
              - 1
            description: |
              0 - AuthenticationRequest: initiate device authentication
              1 - ActivitySubscription: set user activity subscription
          data:
            oneOf:
              - type: 'null'
              - type: object
                required:
                  - users
                properties:
                  users:
                    type: array
                    items:
                      type: string
                      format: uuid
                    description: User UUIDs to subscribe to.
            description: Payload varies by type.
    ServerMessage:
      name: ServerMessage
      contentType: application/json
      payload:
        type: object
        required:
          - id
          - type
          - data
        properties:
          id:
            type: string
            format: uuid
            description: Correlation ID from the originating request.
          type:
            type: integer
            enum:
              - 0
              - 1
              - 2
            description: |
              0 - ActivityData: user activity snapshot
              1 - AuthenticationURL: one-time authentication URL
              2 - Authentication: session token
          data:
            oneOf:
              - type: object
                additionalProperties:
                  type: array
                  items:
                    type: object
                    required:
                      - session_id
                      - track_id
                      - track_name
                      - artists
                      - active
                      - position
                      - start_timestamp
                      - end_timestamp
                      - updated_at
                    properties:
                      session_id:
                        type: string
                        description: Unique identifier for this listening session.
                      track_id:
                        type: string
                        format: uuid
                        description: Track ID.
                      track_name:
                        type: string
                        description: Track title.
                      artists:
                        type: array
                        items:
                          type: object
                          required:
                            - id
                            - name
                          properties:
                            id:
                              type: string
                              format: uuid
                              description: Artist ID.
                            name:
                              type: string
                              description: Artist name.
                        description: Credited track artists.
                      album_name:
                        type: string
                        nullable: true
                        description: Album name.
                      album_asset_url:
                        type: string
                        nullable: true
                        description: Album cover URL.
                      active:
                        type: boolean
                        description: >-
                          Whether the user is currently playing this track
                          (true) or has stopped/paused (false)..
                      position:
                        type: integer
                        format: int64
                        description: Current playback position in seconds.
                      start_timestamp:
                        type: string
                        format: date-time
                        description: When the track was played.
                      end_timestamp:
                        type: string
                        format: date-time
                        description: Projected end time based on track duration.
                      updated_at:
                        type: string
                        format: date-time
                        description: When this activity was last updated.
                description: Map of user UUID to their activities.
              - type: object
                required:
                  - url
                properties:
                  url:
                    type: string
                    format: uri
                    description: One-time URL for the user to authorize the device.
              - type: object
                required:
                  - token
                properties:
                  token:
                    type: string
                    description: JWT authentication token.
            description: Payload varies by type.
  schemas:
    ClientEventType:
      type: integer
      enum:
        - 0
        - 1
      description: |
        0 - AuthenticationRequest: initiate device authentication
        1 - ActivitySubscription: set user activity subscription
    ServerEventType:
      type: integer
      enum:
        - 0
        - 1
        - 2
      description: |
        0 - ActivityData: user activity snapshot
        1 - AuthenticationURL: one-time authentication URL
        2 - Authentication: session token
    ClientMessageStruct:
      type: object
      required:
        - id
        - type
      properties:
        id:
          type: string
          format: uuid
          description: Unique correlation ID. Echoed in server responses.
        type:
          type: integer
          enum:
            - 0
            - 1
          description: |
            0 - AuthenticationRequest: initiate device authentication
            1 - ActivitySubscription: set user activity subscription
        data:
          oneOf:
            - type: 'null'
            - type: object
              required:
                - users
              properties:
                users:
                  type: array
                  items:
                    type: string
                    format: uuid
                  description: User UUIDs to subscribe to.
          description: Payload varies by type.
    ServerMessageStruct:
      type: object
      required:
        - id
        - type
        - data
      properties:
        id:
          type: string
          format: uuid
          description: Correlation ID from the originating request.
        type:
          type: integer
          enum:
            - 0
            - 1
            - 2
          description: |
            0 - ActivityData: user activity snapshot
            1 - AuthenticationURL: one-time authentication URL
            2 - Authentication: session token
        data:
          oneOf:
            - type: object
              additionalProperties:
                type: array
                items:
                  type: object
                  required:
                    - session_id
                    - track_id
                    - track_name
                    - artists
                    - active
                    - position
                    - start_timestamp
                    - end_timestamp
                    - updated_at
                  properties:
                    session_id:
                      type: string
                      description: Unique identifier for this listening session.
                    track_id:
                      type: string
                      format: uuid
                      description: Track ID.
                    track_name:
                      type: string
                      description: Track title.
                    artists:
                      type: array
                      items:
                        type: object
                        required:
                          - id
                          - name
                        properties:
                          id:
                            type: string
                            format: uuid
                            description: Artist ID.
                          name:
                            type: string
                            description: Artist name.
                      description: Credited track artists.
                    album_name:
                      type: string
                      nullable: true
                      description: Album name.
                    album_asset_url:
                      type: string
                      nullable: true
                      description: Album cover URL.
                    active:
                      type: boolean
                      description: >-
                        Whether the user is currently playing this track (true)
                        or has stopped/paused (false)..
                    position:
                      type: integer
                      format: int64
                      description: Current playback position in seconds.
                    start_timestamp:
                      type: string
                      format: date-time
                      description: When the track was played.
                    end_timestamp:
                      type: string
                      format: date-time
                      description: Projected end time based on track duration.
                    updated_at:
                      type: string
                      format: date-time
                      description: When this activity was last updated.
              description: Map of user UUID to their activities.
            - type: object
              required:
                - url
              properties:
                url:
                  type: string
                  format: uri
                  description: One-time URL for the user to authorize the device.
            - type: object
              required:
                - token
              properties:
                token:
                  type: string
                  description: JWT authentication token.
          description: Payload varies by type.
    Subscription:
      type: object
      required:
        - users
      properties:
        users:
          type: array
          items:
            type: string
            format: uuid
          description: User UUIDs to subscribe to.
    ActivityData:
      type: object
      additionalProperties:
        type: array
        items:
          type: object
          required:
            - session_id
            - track_id
            - track_name
            - artists
            - active
            - position
            - start_timestamp
            - end_timestamp
            - updated_at
          properties:
            session_id:
              type: string
              description: Unique identifier for this listening session.
            track_id:
              type: string
              format: uuid
              description: Track ID.
            track_name:
              type: string
              description: Track title.
            artists:
              type: array
              items:
                type: object
                required:
                  - id
                  - name
                properties:
                  id:
                    type: string
                    format: uuid
                    description: Artist ID.
                  name:
                    type: string
                    description: Artist name.
              description: Credited track artists.
            album_name:
              type: string
              nullable: true
              description: Album name.
            album_asset_url:
              type: string
              nullable: true
              description: Album cover URL.
            active:
              type: boolean
              description: >-
                Whether the user is currently playing this track (true) or has
                stopped/paused (false)..
            position:
              type: integer
              format: int64
              description: Current playback position in seconds.
            start_timestamp:
              type: string
              format: date-time
              description: When the track was played.
            end_timestamp:
              type: string
              format: date-time
              description: Projected end time based on track duration.
            updated_at:
              type: string
              format: date-time
              description: When this activity was last updated.
      description: Map of user UUID to their activities.
    AuthURL:
      type: object
      required:
        - url
      properties:
        url:
          type: string
          format: uri
          description: One-time URL for the user to authorize the device.
    AuthToken:
      type: object
      required:
        - token
      properties:
        token:
          type: string
          description: JWT authentication token.
    Activity:
      type: object
      required:
        - session_id
        - track_id
        - track_name
        - artists
        - active
        - position
        - start_timestamp
        - end_timestamp
        - updated_at
      properties:
        session_id:
          type: string
          description: Unique identifier for this listening session.
        track_id:
          type: string
          format: uuid
          description: Track ID.
        track_name:
          type: string
          description: Track title.
        artists:
          type: array
          items:
            type: object
            required:
              - id
              - name
            properties:
              id:
                type: string
                format: uuid
                description: Artist ID.
              name:
                type: string
                description: Artist name.
          description: Credited track artists.
        album_name:
          type: string
          nullable: true
          description: Album name.
        album_asset_url:
          type: string
          nullable: true
          description: Album cover URL.
        active:
          type: boolean
          description: >-
            Whether the user is currently playing this track (true) or has
            stopped/paused (false)..
        position:
          type: integer
          format: int64
          description: Current playback position in seconds.
        start_timestamp:
          type: string
          format: date-time
          description: When the track was played.
        end_timestamp:
          type: string
          format: date-time
          description: Projected end time based on track duration.
        updated_at:
          type: string
          format: date-time
          description: When this activity was last updated.
    Artist:
      type: object
      required:
        - id
        - name
      properties:
        id:
          type: string
          format: uuid
          description: Artist ID.
        name:
          type: string
          description: Artist name.`
