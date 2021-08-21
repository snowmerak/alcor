// Automatically generated rust module for 'auth.proto' file

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
pub struct PrivateKey<'a> {
    pub X: Cow<'a, str>,
    pub Y: Cow<'a, str>,
}

impl<'a> MessageRead<'a> for PrivateKey<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(10) => msg.X = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(18) => msg.Y = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for PrivateKey<'a> {
    fn get_size(&self) -> usize {
        0
        + if self.X == "" { 0 } else { 1 + sizeof_len((&self.X).len()) }
        + if self.Y == "" { 0 } else { 1 + sizeof_len((&self.Y).len()) }
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.X != "" { w.write_with_tag(10, |w| w.write_string(&**&self.X))?; }
        if self.Y != "" { w.write_with_tag(18, |w| w.write_string(&**&self.Y))?; }
        Ok(())
    }
}

#[derive(Debug, Default, PartialEq, Clone)]
pub struct PublicKey<'a> {
    pub X: Cow<'a, str>,
    pub Y: Cow<'a, str>,
}

impl<'a> MessageRead<'a> for PublicKey<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(10) => msg.X = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(18) => msg.Y = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for PublicKey<'a> {
    fn get_size(&self) -> usize {
        0
        + if self.X == "" { 0 } else { 1 + sizeof_len((&self.X).len()) }
        + if self.Y == "" { 0 } else { 1 + sizeof_len((&self.Y).len()) }
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.X != "" { w.write_with_tag(10, |w| w.write_string(&**&self.X))?; }
        if self.Y != "" { w.write_with_tag(18, |w| w.write_string(&**&self.Y))?; }
        Ok(())
    }
}

#[derive(Debug, Default, PartialEq, Clone)]
pub struct Signature<'a> {
    pub R: Cow<'a, str>,
    pub S: Cow<'a, str>,
}

impl<'a> MessageRead<'a> for Signature<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(10) => msg.R = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(18) => msg.S = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for Signature<'a> {
    fn get_size(&self) -> usize {
        0
        + if self.R == "" { 0 } else { 1 + sizeof_len((&self.R).len()) }
        + if self.S == "" { 0 } else { 1 + sizeof_len((&self.S).len()) }
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.R != "" { w.write_with_tag(10, |w| w.write_string(&**&self.R))?; }
        if self.S != "" { w.write_with_tag(18, |w| w.write_string(&**&self.S))?; }
        Ok(())
    }
}

