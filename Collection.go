package Collection

import (
	"math/rand"
)

// Structure for the collection constructor!
type CollectionConstructor struct {
	cache map[string]interface{}
}

// Structure for the basic Dataset!
type Dataset struct {
	key   string
	value interface{}
}

// The main collection function which returns the collection constructor!
func Collection() CollectionConstructor {

	return CollectionConstructor{
		cache: map[string]interface{}{},
	}

}

/**
 * Sets a value for the key!
 */
func (self CollectionConstructor) Set(key string, value interface{}) {
	self.cache[key] = value
}

/**
 * Returns the value of the key but in tupule in the form of (value, found).
 * Found will be a boolean stating is the value found or not.
 * Value will be nil if no id has been registered already!
 */
func (self CollectionConstructor) Get(key string) (value interface{}, found bool) {
	foundvalue, found := self.cache[key]

	if !found {
		return nil, false
	} else {
		return foundvalue, true
	}
}

/**
 * Deletes the value by the key
 */
func (self CollectionConstructor) Delete(key string) {
	delete(self.cache, key)
}

/**
 * Deletes everything what is registered!
 */
func (self CollectionConstructor) DeleteAll() {
	for key := range self.cache {
		delete(self.cache, key)
	}
}

/**
 * Returns an array of datasets containing all the registered keys and values!
 */
func (self CollectionConstructor) All() []Dataset {
	Datasets := []Dataset{}

	for key, value := range self.cache {
		Datasets = append(Datasets, Dataset{key, value})
	}

	return Datasets
}

/**
 * Returns a boolean stating the key is registered or not!
 */
func (self CollectionConstructor) Exists(key string) bool {
	_, found := self.Get(key)
	return found
}

/**
 * Returns an array of strings containing all registered keys!
 */
func (self CollectionConstructor) Keys() []string {
	keys := []string{}

	for key := range self.cache {
		keys = append(keys, key)
	}

	return keys
}

/**
 * Returns an array of registered values!
 */
func (self CollectionConstructor) Values() []interface{} {
	values := []interface{}{}

	for _, value := range self.cache {
		values = append(values, value)
	}

	return values
}

/**
 * Returns a random dataset from CollectionConstructor.All()!
 * If not even 1 key is registered will return "Dataset{ "nil", nil }"
 */
func (self CollectionConstructor) Random() Dataset {
	Datasets := self.All()

	if len(Datasets) == 0 {
		return Dataset{"nil", nil}
	}

	return Datasets[rand.Intn(len(Datasets))]
}

/**
 * Returns the number of keys registered!
 */
func (self CollectionConstructor) Length() int {
	return len(self.cache)
}

/**
 * Runs the callback function with each registered key and value
 */
func (self CollectionConstructor) ForEach(callback func(value interface{}, key string)) {
	for key, value := range self.cache {
		callback(value, key)
	}
}

/**
 * Runs the callback function with each registered key and value.
 * If any dataset passes the test will return that key and value else will return false in the found tupule!
 */
func (self CollectionConstructor) Find(callback func(value interface{}, key string) bool) (value interface{}, found bool) {
	for key, value := range self.cache {
		if callback(value, key) {
			return value, true
		}
	}

	return nil, false
}

/**
 * Runs the callback on every registered key and value.
 * Sends a new collection with the datasets which have passed the callback!
 */
func (self CollectionConstructor) Filter(callback func(value interface{}, key string) bool) CollectionConstructor {
	collection := Collection()

	for key, value := range self.cache {
		if callback(value, key) {
			collection.Set(key, value)
		}
	}

	return collection
}

/**
 * Runs the callback on every registered key and value.
 * If it fails the callback, will delete it from the cache!
 */
func (self CollectionConstructor) Sweep(callback func(value interface{}, key string) bool) {
	for key, value := range self.cache {
		if !callback(value, key) {
			delete(self.cache, key)
		}
	}
}

/**
 * Runs the callback on every registered key and value.
 * Replaces the values of each key with the value returned by the callback!
 */
func (self CollectionConstructor) Map(callback func(value interface{}, key string) interface{}) CollectionConstructor {
	collection := Collection()

	for key, value := range self.cache {
		newvalue := callback(value, key)
		collection.Set(key, newvalue)
	}

	return collection
}

/**
 * Adds all the keys and values of other collections too!
 */
func (self CollectionConstructor) Concat(collections ...CollectionConstructor) {
	for _, collection := range collections {
		for key, value := range collection.cache {
			self.cache[key] = value
		}
	}
}
