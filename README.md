# livesort

This is an API for dynamic sorting. It is dynamic because it allows data to be added or removed during sorting. It also makes it possible to stop sorting early and get an "approximately sorted" list.

This implementation is not designed to be computationally-efficient. Rather, it is intended to sort lists with relatively few entries when comparisons are extremely costly (e.g. a human has to perform them). Using this API to fully sort lists (even lists with only a few thousand entries) would take a very long time.
