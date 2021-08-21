// Automatically generated rust module for 'capsule.proto' file

#![allow(non_snake_case)]
#![allow(non_upper_case_globals)]
#![allow(non_camel_case_types)]
#![allow(unused_imports)]
#![allow(unknown_lints)]
#![allow(clippy::all)]
#![cfg_attr(rustfmt, rustfmt_skip)]


use std::borrow::Cow;
use quick_protobuf::{MessageRead, MessageWrite, BytesReader, Writer, WriterBackend, Result};
use quick_protobuf::sizeofs::*;
use super::*;

#[derive(Debug, Default, PartialEq, Clone)]
pub struct Capsule<'a> {
    pub ID: Cow<'a, str>,
    pub Data: Cow<'a, [u8]>,
}

impl<'a> MessageRead<'a> for Capsule<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(10) => msg.ID = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(18) => msg.Data = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for Capsule<'a> {
    fn get_size(&self) -> usize {
        0
        + if self.ID == "" { 0 } else { 1 + sizeof_len((&self.ID).len()) }
        + if self.Data == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.Data).len()) }
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.ID != "" { w.write_with_tag(10, |w| w.write_string(&**&self.ID))?; }
        if self.Data != Cow::Borrowed(b"") { w.write_with_tag(18, |w| w.write_bytes(&**&self.Data))?; }
        Ok(())
    }
}

