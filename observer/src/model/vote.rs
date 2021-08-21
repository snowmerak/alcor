// Automatically generated rust module for 'vote.proto' file

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
pub struct KeyBytes<'a> {
    pub Data: Cow<'a, [u8]>,
}

impl<'a> MessageRead<'a> for KeyBytes<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(10) => msg.Data = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for KeyBytes<'a> {
    fn get_size(&self) -> usize {
        0
        + if self.Data == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.Data).len()) }
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.Data != Cow::Borrowed(b"") { w.write_with_tag(10, |w| w.write_bytes(&**&self.Data))?; }
        Ok(())
    }
}

#[derive(Debug, Default, PartialEq, Clone)]
pub struct Paper<'a> {
    pub VoterId: Cow<'a, [u8]>,
    pub Timestamp: Cow<'a, [u8]>,
    pub Message: Cow<'a, [u8]>,
    pub RandomBytes: Cow<'a, [u8]>,
    pub Hash: Cow<'a, [u8]>,
    pub Signature: Cow<'a, [u8]>,
}

impl<'a> MessageRead<'a> for Paper<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(18) => msg.VoterId = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(26) => msg.Timestamp = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(34) => msg.Message = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(42) => msg.RandomBytes = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(50) => msg.Hash = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(58) => msg.Signature = r.read_bytes(bytes).map(Cow::Borrowed)?,
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for Paper<'a> {
    fn get_size(&self) -> usize {
        0
        + if self.VoterId == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.VoterId).len()) }
        + if self.Timestamp == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.Timestamp).len()) }
        + if self.Message == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.Message).len()) }
        + if self.RandomBytes == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.RandomBytes).len()) }
        + if self.Hash == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.Hash).len()) }
        + if self.Signature == Cow::Borrowed(b"") { 0 } else { 1 + sizeof_len((&self.Signature).len()) }
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.VoterId != Cow::Borrowed(b"") { w.write_with_tag(18, |w| w.write_bytes(&**&self.VoterId))?; }
        if self.Timestamp != Cow::Borrowed(b"") { w.write_with_tag(26, |w| w.write_bytes(&**&self.Timestamp))?; }
        if self.Message != Cow::Borrowed(b"") { w.write_with_tag(34, |w| w.write_bytes(&**&self.Message))?; }
        if self.RandomBytes != Cow::Borrowed(b"") { w.write_with_tag(42, |w| w.write_bytes(&**&self.RandomBytes))?; }
        if self.Hash != Cow::Borrowed(b"") { w.write_with_tag(50, |w| w.write_bytes(&**&self.Hash))?; }
        if self.Signature != Cow::Borrowed(b"") { w.write_with_tag(58, |w| w.write_bytes(&**&self.Signature))?; }
        Ok(())
    }
}

#[derive(Debug, Default, PartialEq, Clone)]
pub struct Result_pb<'a> {
    pub Ok_pb: bool,
    pub Error: Cow<'a, str>,
}

impl<'a> MessageRead<'a> for Result_pb<'a> {
    fn from_reader(r: &mut BytesReader, bytes: &'a [u8]) -> Result<Self> {
        let mut msg = Self::default();
        while !r.is_eof() {
            match r.next_tag(bytes) {
                Ok(8) => msg.Ok_pb = r.read_bool(bytes)?,
                Ok(18) => msg.Error = r.read_string(bytes).map(Cow::Borrowed)?,
                Ok(t) => { r.read_unknown(bytes, t)?; }
                Err(e) => return Err(e),
            }
        }
        Ok(msg)
    }
}

impl<'a> MessageWrite for Result_pb<'a> {
    fn get_size(&self) -> usize {
        0
        + if self.Ok_pb == false { 0 } else { 1 + sizeof_varint(*(&self.Ok_pb) as u64) }
        + if self.Error == "" { 0 } else { 1 + sizeof_len((&self.Error).len()) }
    }

    fn write_message<W: WriterBackend>(&self, w: &mut Writer<W>) -> Result<()> {
        if self.Ok_pb != false { w.write_with_tag(8, |w| w.write_bool(*&self.Ok_pb))?; }
        if self.Error != "" { w.write_with_tag(18, |w| w.write_string(&**&self.Error))?; }
        Ok(())
    }
}

