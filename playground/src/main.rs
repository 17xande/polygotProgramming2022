fn error_me(throw: bool) -> Result<(), usize> {
    if throw {
        return Err(7);
    }

    return Ok(());
}

fn main() -> Result<(), usize> {
    let value = match error_me(true) {
        Err(e) => return Err(e),
        Ok(v) => v,
    };

    // The above can also be achieved with this shorthand:
    let value = error_me(true)?;
    // This allows us to get rid of the "if err" construct thing.
    // The limitation here is that this function needs to return the same
    // thing as the error_me function.

    // Another way of doing it:
    if error_me(false).is_ok() {
        // do something
    }

    // Also look at .unwrap_or(); .expext(); .unwrap()
    
    return Ok(value)
}
