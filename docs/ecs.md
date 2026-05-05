# Go ECS Framework

An entity is an ID. Entities can then have components associated to them. Components are a structured set of fields that reprisent application state. Component fields are mutated by Systems. Systems are the application logic. Systems manage the applications state by mutating an entities with the required components.

### Core logical flow

The system must be able to create entities, associate components to them, and expose them so systems can mutate them efficiently.

# Entities

An id with the ability to store components

# Components

# Systems

# Questions

- How many registries do we need?
  3, Entity registry host state, Component registry host entities with that component, Systems are just the systems.

  The Component Registrey is essentally just a ComponentEntityMap

- What do Systems care about?
  Does a system care about entities? Indirectly yes. They are the logical association between components. If the System requires the PositionComponent & InputComponent, they need to know those 2 components are associated in 1 entity.

  How can a system efficiently query for entities with the correct component associations.

  getEntitiesByComponents(Components...) -> its a filtered list. First component gets all the entities that are associated to it, for the next component only keep entities that are in both .... continue. You end up with a subset of entities that have all the required components. Fetch and return those entities.
  or
  Order Independent XOr sum hash. -> no this would require all permutations to be hashed independently.
