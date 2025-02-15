# Doc

Main
- App
    - world *ecs.World
    - schedule []System
        - ScheduleLabel
            - LOAD
            - UPDATE
            - DRAW
            - EXIT
        - System(world *ecs.World)

functions you can use
App
    -GetSystems
    -SetSystems
    -AddSystems
