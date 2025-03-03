use anyhow::Result;
use starlark::collections::SmallMap;
use starlark::values::dict::Dict;
use starlark::values::Heap;
use std::env;

pub fn get_env(starlark_heap: &Heap) -> Result<Dict> {
    let res = SmallMap::new();
    let mut dict_res = Dict::new(res);

    for (key, val) in env::vars() {
        let key_value = starlark_heap.alloc_str(&key);
        let val_value = starlark_heap.alloc_str(&val);
        dict_res.insert_hashed(
            match key_value.to_value().get_hashed() {
                Ok(val) => val,
                Err(e) => return Err(anyhow::anyhow!("Failed to alloc name information: {}", e)),
            },
            val_value.to_value(),
        );
    }
    
    Ok(dict_res)
}

#[cfg(test)]
mod tests {
    use starlark::{values::Value, const_frozen_string};
    use anyhow::Result;
    use std::env;
    use super::*;

    #[test]
    fn test_get_env() -> Result<()> {
        env::set_var("FOO", "BAR");
        let test_heap = Heap::new();
        let res = get_env(&test_heap)?;
        let val: Value<'_> = res.get(const_frozen_string!("FOO").to_value())?.unwrap();
        assert_eq!(val.unpack_str().unwrap(), "BAR");
        Ok(())
    }
}
