// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

(function (root, factory) {
  if (typeof define === 'function' && define.amd) {
    define(['kaitai-struct/KaitaiStream'], factory);
  } else if (typeof module === 'object' && module.exports) {
    module.exports = factory(require('kaitai-struct/KaitaiStream'));
  } else {
    root.Elf = factory(root.KaitaiStream);
  }
}(this, function (KaitaiStream) {
/**
 * @see {@link https://sourceware.org/git/?p=glibc.git;a=blob;f=elf/elf.h;hb=HEAD|Source}
 */

var Elf = (function() {
  Elf.Endian = Object.freeze({
    LE: 1,
    BE: 2,

    1: "LE",
    2: "BE",
  });

  Elf.ShType = Object.freeze({
    NULL_TYPE: 0,
    PROGBITS: 1,
    SYMTAB: 2,
    STRTAB: 3,
    RELA: 4,
    HASH: 5,
    DYNAMIC: 6,
    NOTE: 7,
    NOBITS: 8,
    REL: 9,
    SHLIB: 10,
    DYNSYM: 11,
    INIT_ARRAY: 14,
    FINI_ARRAY: 15,
    PREINIT_ARRAY: 16,
    GROUP: 17,
    SYMTAB_SHNDX: 18,
    SUNW_CAPCHAIN: 1879048175,
    SUNW_CAPINFO: 1879048176,
    SUNW_SYMSORT: 1879048177,
    SUNW_TLSSORT: 1879048178,
    SUNW_LDYNSYM: 1879048179,
    SUNW_DOF: 1879048180,
    SUNW_CAP: 1879048181,
    SUNW_SIGNATURE: 1879048182,
    SUNW_ANNOTATE: 1879048183,
    SUNW_DEBUGSTR: 1879048184,
    SUNW_DEBUG: 1879048185,
    SUNW_MOVE: 1879048186,
    SUNW_COMDAT: 1879048187,
    SUNW_SYMINFO: 1879048188,
    SUNW_VERDEF: 1879048189,
    SUNW_VERNEED: 1879048190,
    SUNW_VERSYM: 1879048191,
    SPARC_GOTDATA: 1879048192,
    ARM_EXIDX: 1879048193,
    ARM_PREEMPTMAP: 1879048194,
    ARM_ATTRIBUTES: 1879048195,

    0: "NULL_TYPE",
    1: "PROGBITS",
    2: "SYMTAB",
    3: "STRTAB",
    4: "RELA",
    5: "HASH",
    6: "DYNAMIC",
    7: "NOTE",
    8: "NOBITS",
    9: "REL",
    10: "SHLIB",
    11: "DYNSYM",
    14: "INIT_ARRAY",
    15: "FINI_ARRAY",
    16: "PREINIT_ARRAY",
    17: "GROUP",
    18: "SYMTAB_SHNDX",
    1879048175: "SUNW_CAPCHAIN",
    1879048176: "SUNW_CAPINFO",
    1879048177: "SUNW_SYMSORT",
    1879048178: "SUNW_TLSSORT",
    1879048179: "SUNW_LDYNSYM",
    1879048180: "SUNW_DOF",
    1879048181: "SUNW_CAP",
    1879048182: "SUNW_SIGNATURE",
    1879048183: "SUNW_ANNOTATE",
    1879048184: "SUNW_DEBUGSTR",
    1879048185: "SUNW_DEBUG",
    1879048186: "SUNW_MOVE",
    1879048187: "SUNW_COMDAT",
    1879048188: "SUNW_SYMINFO",
    1879048189: "SUNW_VERDEF",
    1879048190: "SUNW_VERNEED",
    1879048191: "SUNW_VERSYM",
    1879048192: "SPARC_GOTDATA",
    1879048193: "ARM_EXIDX",
    1879048194: "ARM_PREEMPTMAP",
    1879048195: "ARM_ATTRIBUTES",
  });

  Elf.OsAbi = Object.freeze({
    SYSTEM_V: 0,
    HP_UX: 1,
    NETBSD: 2,
    GNU: 3,
    SOLARIS: 6,
    AIX: 7,
    IRIX: 8,
    FREEBSD: 9,
    TRU64: 10,
    MODESTO: 11,
    OPENBSD: 12,
    OPENVMS: 13,
    NSK: 14,
    AROS: 15,
    FENIXOS: 16,
    CLOUDABI: 17,
    OPENVOS: 18,

    0: "SYSTEM_V",
    1: "HP_UX",
    2: "NETBSD",
    3: "GNU",
    6: "SOLARIS",
    7: "AIX",
    8: "IRIX",
    9: "FREEBSD",
    10: "TRU64",
    11: "MODESTO",
    12: "OPENBSD",
    13: "OPENVMS",
    14: "NSK",
    15: "AROS",
    16: "FENIXOS",
    17: "CLOUDABI",
    18: "OPENVOS",
  });

  Elf.Machine = Object.freeze({
    NOT_SET: 0,
    SPARC: 2,
    X86: 3,
    MIPS: 8,
    POWERPC: 20,
    ARM: 40,
    SUPERH: 42,
    IA_64: 50,
    X86_64: 62,
    AARCH64: 183,
    RISCV: 243,
    BPF: 247,

    0: "NOT_SET",
    2: "SPARC",
    3: "X86",
    8: "MIPS",
    20: "POWERPC",
    40: "ARM",
    42: "SUPERH",
    50: "IA_64",
    62: "X86_64",
    183: "AARCH64",
    243: "RISCV",
    247: "BPF",
  });

  Elf.DynamicArrayTags = Object.freeze({
    NULL: 0,
    NEEDED: 1,
    PLTRELSZ: 2,
    PLTGOT: 3,
    HASH: 4,
    STRTAB: 5,
    SYMTAB: 6,
    RELA: 7,
    RELASZ: 8,
    RELAENT: 9,
    STRSZ: 10,
    SYMENT: 11,
    INIT: 12,
    FINI: 13,
    SONAME: 14,
    RPATH: 15,
    SYMBOLIC: 16,
    REL: 17,
    RELSZ: 18,
    RELENT: 19,
    PLTREL: 20,
    DEBUG: 21,
    TEXTREL: 22,
    JMPREL: 23,
    BIND_NOW: 24,
    INIT_ARRAY: 25,
    FINI_ARRAY: 26,
    INIT_ARRAYSZ: 27,
    FINI_ARRAYSZ: 28,
    RUNPATH: 29,
    FLAGS: 30,
    PREINIT_ARRAY: 32,
    PREINIT_ARRAYSZ: 33,
    MAXPOSTAGS: 34,
    SUNW_AUXILIARY: 1610612749,
    SUNW_FILTER: 1610612750,
    SUNW_CAP: 1610612752,
    SUNW_SYMTAB: 1610612753,
    SUNW_SYMSZ: 1610612754,
    SUNW_SORTENT: 1610612755,
    SUNW_SYMSORT: 1610612756,
    SUNW_SYMSORTSZ: 1610612757,
    SUNW_TLSSORT: 1610612758,
    SUNW_TLSSORTSZ: 1610612759,
    SUNW_CAPINFO: 1610612760,
    SUNW_STRPAD: 1610612761,
    SUNW_CAPCHAIN: 1610612762,
    SUNW_LDMACH: 1610612763,
    SUNW_CAPCHAINENT: 1610612765,
    SUNW_CAPCHAINSZ: 1610612767,
    HIOS: 1879044096,
    VALRNGLO: 1879047424,
    GNU_PRELINKED: 1879047669,
    GNU_CONFLICTSZ: 1879047670,
    GNU_LIBLISTSZ: 1879047671,
    CHECKSUM: 1879047672,
    PLTPADSZ: 1879047673,
    MOVEENT: 1879047674,
    MOVESZ: 1879047675,
    FEATURE_1: 1879047676,
    POSFLAG_1: 1879047677,
    SYMINSZ: 1879047678,
    VALRNGHI: 1879047679,
    ADDRRNGLO: 1879047680,
    GNU_HASH: 1879047925,
    TLSDESC_PLT: 1879047926,
    TLSDESC_GOT: 1879047927,
    GNU_CONFLICT: 1879047928,
    GNU_LIBLIST: 1879047929,
    CONFIG: 1879047930,
    DEPAUDIT: 1879047931,
    AUDIT: 1879047932,
    PLTPAD: 1879047933,
    MOVETAB: 1879047934,
    ADDRRNGHI: 1879047935,
    VERSYM: 1879048176,
    RELACOUNT: 1879048185,
    RELCOUNT: 1879048186,
    FLAGS_1: 1879048187,
    VERDEF: 1879048188,
    VERDEFNUM: 1879048189,
    VERNEED: 1879048190,
    VERNEEDNUM: 1879048191,
    LOPROC: 1879048192,
    SPARC_REGISTER: 1879048193,
    AUXILIARY: 2147483645,
    USED: 2147483646,
    HIPROC: 2147483647,

    0: "NULL",
    1: "NEEDED",
    2: "PLTRELSZ",
    3: "PLTGOT",
    4: "HASH",
    5: "STRTAB",
    6: "SYMTAB",
    7: "RELA",
    8: "RELASZ",
    9: "RELAENT",
    10: "STRSZ",
    11: "SYMENT",
    12: "INIT",
    13: "FINI",
    14: "SONAME",
    15: "RPATH",
    16: "SYMBOLIC",
    17: "REL",
    18: "RELSZ",
    19: "RELENT",
    20: "PLTREL",
    21: "DEBUG",
    22: "TEXTREL",
    23: "JMPREL",
    24: "BIND_NOW",
    25: "INIT_ARRAY",
    26: "FINI_ARRAY",
    27: "INIT_ARRAYSZ",
    28: "FINI_ARRAYSZ",
    29: "RUNPATH",
    30: "FLAGS",
    32: "PREINIT_ARRAY",
    33: "PREINIT_ARRAYSZ",
    34: "MAXPOSTAGS",
    1610612749: "SUNW_AUXILIARY",
    1610612750: "SUNW_FILTER",
    1610612752: "SUNW_CAP",
    1610612753: "SUNW_SYMTAB",
    1610612754: "SUNW_SYMSZ",
    1610612755: "SUNW_SORTENT",
    1610612756: "SUNW_SYMSORT",
    1610612757: "SUNW_SYMSORTSZ",
    1610612758: "SUNW_TLSSORT",
    1610612759: "SUNW_TLSSORTSZ",
    1610612760: "SUNW_CAPINFO",
    1610612761: "SUNW_STRPAD",
    1610612762: "SUNW_CAPCHAIN",
    1610612763: "SUNW_LDMACH",
    1610612765: "SUNW_CAPCHAINENT",
    1610612767: "SUNW_CAPCHAINSZ",
    1879044096: "HIOS",
    1879047424: "VALRNGLO",
    1879047669: "GNU_PRELINKED",
    1879047670: "GNU_CONFLICTSZ",
    1879047671: "GNU_LIBLISTSZ",
    1879047672: "CHECKSUM",
    1879047673: "PLTPADSZ",
    1879047674: "MOVEENT",
    1879047675: "MOVESZ",
    1879047676: "FEATURE_1",
    1879047677: "POSFLAG_1",
    1879047678: "SYMINSZ",
    1879047679: "VALRNGHI",
    1879047680: "ADDRRNGLO",
    1879047925: "GNU_HASH",
    1879047926: "TLSDESC_PLT",
    1879047927: "TLSDESC_GOT",
    1879047928: "GNU_CONFLICT",
    1879047929: "GNU_LIBLIST",
    1879047930: "CONFIG",
    1879047931: "DEPAUDIT",
    1879047932: "AUDIT",
    1879047933: "PLTPAD",
    1879047934: "MOVETAB",
    1879047935: "ADDRRNGHI",
    1879048176: "VERSYM",
    1879048185: "RELACOUNT",
    1879048186: "RELCOUNT",
    1879048187: "FLAGS_1",
    1879048188: "VERDEF",
    1879048189: "VERDEFNUM",
    1879048190: "VERNEED",
    1879048191: "VERNEEDNUM",
    1879048192: "LOPROC",
    1879048193: "SPARC_REGISTER",
    2147483645: "AUXILIARY",
    2147483646: "USED",
    2147483647: "HIPROC",
  });

  Elf.Bits = Object.freeze({
    B32: 1,
    B64: 2,

    1: "B32",
    2: "B64",
  });

  Elf.PhType = Object.freeze({
    NULL_TYPE: 0,
    LOAD: 1,
    DYNAMIC: 2,
    INTERP: 3,
    NOTE: 4,
    SHLIB: 5,
    PHDR: 6,
    TLS: 7,
    GNU_EH_FRAME: 1685382480,
    GNU_STACK: 1685382481,
    GNU_RELRO: 1685382482,
    PAX_FLAGS: 1694766464,
    HIOS: 1879048191,
    ARM_EXIDX: 1879048193,

    0: "NULL_TYPE",
    1: "LOAD",
    2: "DYNAMIC",
    3: "INTERP",
    4: "NOTE",
    5: "SHLIB",
    6: "PHDR",
    7: "TLS",
    1685382480: "GNU_EH_FRAME",
    1685382481: "GNU_STACK",
    1685382482: "GNU_RELRO",
    1694766464: "PAX_FLAGS",
    1879048191: "HIOS",
    1879048193: "ARM_EXIDX",
  });

  Elf.ObjType = Object.freeze({
    RELOCATABLE: 1,
    EXECUTABLE: 2,
    SHARED: 3,
    CORE: 4,

    1: "RELOCATABLE",
    2: "EXECUTABLE",
    3: "SHARED",
    4: "CORE",
  });

  function Elf(_io, _parent, _root) {
    this._io = _io;
    this._parent = _parent;
    this._root = _root || this;

    this._read();
  }
  Elf.prototype._read = function() {
    this.magic = this._io.ensureFixedContents([127, 69, 76, 70]);
    this.bits = this._io.readU1();
    this.endian = this._io.readU1();
    this.eiVersion = this._io.readU1();
    this.abi = this._io.readU1();
    this.abiVersion = this._io.readU1();
    this.pad = this._io.readBytes(7);
    this.header = new EndianElf(this._io, this, this._root);
  }

  var PhdrTypeFlags = Elf.PhdrTypeFlags = (function() {
    function PhdrTypeFlags(_io, _parent, _root, value) {
      this._io = _io;
      this._parent = _parent;
      this._root = _root || this;
      this.value = value;

      this._read();
    }
    PhdrTypeFlags.prototype._read = function() {
    }
    Object.defineProperty(PhdrTypeFlags.prototype, 'read', {
      get: function() {
        if (this._m_read !== undefined)
          return this._m_read;
        this._m_read = (this.value & 4) != 0;
        return this._m_read;
      }
    });
    Object.defineProperty(PhdrTypeFlags.prototype, 'write', {
      get: function() {
        if (this._m_write !== undefined)
          return this._m_write;
        this._m_write = (this.value & 2) != 0;
        return this._m_write;
      }
    });
    Object.defineProperty(PhdrTypeFlags.prototype, 'execute', {
      get: function() {
        if (this._m_execute !== undefined)
          return this._m_execute;
        this._m_execute = (this.value & 1) != 0;
        return this._m_execute;
      }
    });
    Object.defineProperty(PhdrTypeFlags.prototype, 'maskProc', {
      get: function() {
        if (this._m_maskProc !== undefined)
          return this._m_maskProc;
        this._m_maskProc = (this.value & 4026531840) != 0;
        return this._m_maskProc;
      }
    });

    return PhdrTypeFlags;
  })();

  var SectionHeaderFlags = Elf.SectionHeaderFlags = (function() {
    function SectionHeaderFlags(_io, _parent, _root, value) {
      this._io = _io;
      this._parent = _parent;
      this._root = _root || this;
      this.value = value;

      this._read();
    }
    SectionHeaderFlags.prototype._read = function() {
    }

    /**
     * might be merged
     */
    Object.defineProperty(SectionHeaderFlags.prototype, 'merge', {
      get: function() {
        if (this._m_merge !== undefined)
          return this._m_merge;
        this._m_merge = (this.value & 16) != 0;
        return this._m_merge;
      }
    });

    /**
     * OS-specific
     */
    Object.defineProperty(SectionHeaderFlags.prototype, 'maskOs', {
      get: function() {
        if (this._m_maskOs !== undefined)
          return this._m_maskOs;
        this._m_maskOs = (this.value & 267386880) != 0;
        return this._m_maskOs;
      }
    });

    /**
     * section is excluded unless referenced or allocated (Solaris)
     */
    Object.defineProperty(SectionHeaderFlags.prototype, 'exclude', {
      get: function() {
        if (this._m_exclude !== undefined)
          return this._m_exclude;
        this._m_exclude = (this.value & 134217728) != 0;
        return this._m_exclude;
      }
    });

    /**
     * Processor-specific
     */
    Object.defineProperty(SectionHeaderFlags.prototype, 'maskProc', {
      get: function() {
        if (this._m_maskProc !== undefined)
          return this._m_maskProc;
        this._m_maskProc = (this.value & 4026531840) != 0;
        return this._m_maskProc;
      }
    });

    /**
     * contains nul-terminated strings
     */
    Object.defineProperty(SectionHeaderFlags.prototype, 'strings', {
      get: function() {
        if (this._m_strings !== undefined)
          return this._m_strings;
        this._m_strings = (this.value & 32) != 0;
        return this._m_strings;
      }
    });

    /**
     * non-standard OS specific handling required
     */
    Object.defineProperty(SectionHeaderFlags.prototype, 'osNonConforming', {
      get: function() {
        if (this._m_osNonConforming !== undefined)
          return this._m_osNonConforming;
        this._m_osNonConforming = (this.value & 256) != 0;
        return this._m_osNonConforming;
      }
    });

    /**
     * occupies memory during execution
     */
    Object.defineProperty(SectionHeaderFlags.prototype, 'alloc', {
      get: function() {
        if (this._m_alloc !== undefined)
          return this._m_alloc;
        this._m_alloc = (this.value & 2) != 0;
        return this._m_alloc;
      }
    });

    /**
     * executable
     */
    Object.defineProperty(SectionHeaderFlags.prototype, 'execInstr', {
      get: function() {
        if (this._m_execInstr !== undefined)
          return this._m_execInstr;
        this._m_execInstr = (this.value & 4) != 0;
        return this._m_execInstr;
      }
    });

    /**
     * 'sh_info' contains SHT index
     */
    Object.defineProperty(SectionHeaderFlags.prototype, 'infoLink', {
      get: function() {
        if (this._m_infoLink !== undefined)
          return this._m_infoLink;
        this._m_infoLink = (this.value & 64) != 0;
        return this._m_infoLink;
      }
    });

    /**
     * writable
     */
    Object.defineProperty(SectionHeaderFlags.prototype, 'write', {
      get: function() {
        if (this._m_write !== undefined)
          return this._m_write;
        this._m_write = (this.value & 1) != 0;
        return this._m_write;
      }
    });

    /**
     * preserve order after combining
     */
    Object.defineProperty(SectionHeaderFlags.prototype, 'linkOrder', {
      get: function() {
        if (this._m_linkOrder !== undefined)
          return this._m_linkOrder;
        this._m_linkOrder = (this.value & 128) != 0;
        return this._m_linkOrder;
      }
    });

    /**
     * special ordering requirement (Solaris)
     */
    Object.defineProperty(SectionHeaderFlags.prototype, 'ordered', {
      get: function() {
        if (this._m_ordered !== undefined)
          return this._m_ordered;
        this._m_ordered = (this.value & 67108864) != 0;
        return this._m_ordered;
      }
    });

    /**
     * section hold thread-local data
     */
    Object.defineProperty(SectionHeaderFlags.prototype, 'tls', {
      get: function() {
        if (this._m_tls !== undefined)
          return this._m_tls;
        this._m_tls = (this.value & 1024) != 0;
        return this._m_tls;
      }
    });

    /**
     * section is member of a group
     */
    Object.defineProperty(SectionHeaderFlags.prototype, 'group', {
      get: function() {
        if (this._m_group !== undefined)
          return this._m_group;
        this._m_group = (this.value & 512) != 0;
        return this._m_group;
      }
    });

    return SectionHeaderFlags;
  })();

  var DtFlag1Values = Elf.DtFlag1Values = (function() {
    function DtFlag1Values(_io, _parent, _root, value) {
      this._io = _io;
      this._parent = _parent;
      this._root = _root || this;
      this.value = value;

      this._read();
    }
    DtFlag1Values.prototype._read = function() {
    }

    /**
     * Singleton symbols are used.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'singleton', {
      get: function() {
        if (this._m_singleton !== undefined)
          return this._m_singleton;
        this._m_singleton = (this.value & 33554432) != 0;
        return this._m_singleton;
      }
    });
    Object.defineProperty(DtFlag1Values.prototype, 'ignmuldef', {
      get: function() {
        if (this._m_ignmuldef !== undefined)
          return this._m_ignmuldef;
        this._m_ignmuldef = (this.value & 262144) != 0;
        return this._m_ignmuldef;
      }
    });

    /**
     * Trigger filtee loading at runtime.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'loadfltr', {
      get: function() {
        if (this._m_loadfltr !== undefined)
          return this._m_loadfltr;
        this._m_loadfltr = (this.value & 16) != 0;
        return this._m_loadfltr;
      }
    });

    /**
     * Set RTLD_INITFIRST for this object
     */
    Object.defineProperty(DtFlag1Values.prototype, 'initfirst', {
      get: function() {
        if (this._m_initfirst !== undefined)
          return this._m_initfirst;
        this._m_initfirst = (this.value & 32) != 0;
        return this._m_initfirst;
      }
    });

    /**
     * Object has individual interposers.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'symintpose', {
      get: function() {
        if (this._m_symintpose !== undefined)
          return this._m_symintpose;
        this._m_symintpose = (this.value & 8388608) != 0;
        return this._m_symintpose;
      }
    });
    Object.defineProperty(DtFlag1Values.prototype, 'noreloc', {
      get: function() {
        if (this._m_noreloc !== undefined)
          return this._m_noreloc;
        this._m_noreloc = (this.value & 4194304) != 0;
        return this._m_noreloc;
      }
    });

    /**
     * Configuration alternative created.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'confalt', {
      get: function() {
        if (this._m_confalt !== undefined)
          return this._m_confalt;
        this._m_confalt = (this.value & 8192) != 0;
        return this._m_confalt;
      }
    });

    /**
     * Disp reloc applied at build time.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'dispreldne', {
      get: function() {
        if (this._m_dispreldne !== undefined)
          return this._m_dispreldne;
        this._m_dispreldne = (this.value & 32768) != 0;
        return this._m_dispreldne;
      }
    });

    /**
     * Set RTLD_GLOBAL for this object.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'rtldGlobal', {
      get: function() {
        if (this._m_rtldGlobal !== undefined)
          return this._m_rtldGlobal;
        this._m_rtldGlobal = (this.value & 2) != 0;
        return this._m_rtldGlobal;
      }
    });

    /**
     * Set RTLD_NODELETE for this object.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'nodelete', {
      get: function() {
        if (this._m_nodelete !== undefined)
          return this._m_nodelete;
        this._m_nodelete = (this.value & 8) != 0;
        return this._m_nodelete;
      }
    });
    Object.defineProperty(DtFlag1Values.prototype, 'trans', {
      get: function() {
        if (this._m_trans !== undefined)
          return this._m_trans;
        this._m_trans = (this.value & 512) != 0;
        return this._m_trans;
      }
    });

    /**
     * $ORIGIN must be handled.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'origin', {
      get: function() {
        if (this._m_origin !== undefined)
          return this._m_origin;
        this._m_origin = (this.value & 128) != 0;
        return this._m_origin;
      }
    });

    /**
     * Set RTLD_NOW for this object.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'now', {
      get: function() {
        if (this._m_now !== undefined)
          return this._m_now;
        this._m_now = (this.value & 1) != 0;
        return this._m_now;
      }
    });
    Object.defineProperty(DtFlag1Values.prototype, 'nohdr', {
      get: function() {
        if (this._m_nohdr !== undefined)
          return this._m_nohdr;
        this._m_nohdr = (this.value & 1048576) != 0;
        return this._m_nohdr;
      }
    });

    /**
     * Filtee terminates filters search.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'endfiltee', {
      get: function() {
        if (this._m_endfiltee !== undefined)
          return this._m_endfiltee;
        this._m_endfiltee = (this.value & 16384) != 0;
        return this._m_endfiltee;
      }
    });

    /**
     * Object has no-direct binding.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'nodirect', {
      get: function() {
        if (this._m_nodirect !== undefined)
          return this._m_nodirect;
        this._m_nodirect = (this.value & 131072) != 0;
        return this._m_nodirect;
      }
    });

    /**
     * Global auditing required.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'globaudit', {
      get: function() {
        if (this._m_globaudit !== undefined)
          return this._m_globaudit;
        this._m_globaudit = (this.value & 16777216) != 0;
        return this._m_globaudit;
      }
    });
    Object.defineProperty(DtFlag1Values.prototype, 'noksyms', {
      get: function() {
        if (this._m_noksyms !== undefined)
          return this._m_noksyms;
        this._m_noksyms = (this.value & 524288) != 0;
        return this._m_noksyms;
      }
    });

    /**
     * Object is used to interpose.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'interpose', {
      get: function() {
        if (this._m_interpose !== undefined)
          return this._m_interpose;
        this._m_interpose = (this.value & 1024) != 0;
        return this._m_interpose;
      }
    });

    /**
     * Object can't be dldump'ed.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'nodump', {
      get: function() {
        if (this._m_nodump !== undefined)
          return this._m_nodump;
        this._m_nodump = (this.value & 4096) != 0;
        return this._m_nodump;
      }
    });

    /**
     * Disp reloc applied at run-time.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'disprelpnd', {
      get: function() {
        if (this._m_disprelpnd !== undefined)
          return this._m_disprelpnd;
        this._m_disprelpnd = (this.value & 65536) != 0;
        return this._m_disprelpnd;
      }
    });

    /**
     * Set RTLD_NOOPEN for this object.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'noopen', {
      get: function() {
        if (this._m_noopen !== undefined)
          return this._m_noopen;
        this._m_noopen = (this.value & 64) != 0;
        return this._m_noopen;
      }
    });
    Object.defineProperty(DtFlag1Values.prototype, 'stub', {
      get: function() {
        if (this._m_stub !== undefined)
          return this._m_stub;
        this._m_stub = (this.value & 67108864) != 0;
        return this._m_stub;
      }
    });

    /**
     * Direct binding enabled.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'direct', {
      get: function() {
        if (this._m_direct !== undefined)
          return this._m_direct;
        this._m_direct = (this.value & 256) != 0;
        return this._m_direct;
      }
    });

    /**
     * Object is modified after built.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'edited', {
      get: function() {
        if (this._m_edited !== undefined)
          return this._m_edited;
        this._m_edited = (this.value & 2097152) != 0;
        return this._m_edited;
      }
    });

    /**
     * Set RTLD_GROUP for this object.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'group', {
      get: function() {
        if (this._m_group !== undefined)
          return this._m_group;
        this._m_group = (this.value & 4) != 0;
        return this._m_group;
      }
    });
    Object.defineProperty(DtFlag1Values.prototype, 'pie', {
      get: function() {
        if (this._m_pie !== undefined)
          return this._m_pie;
        this._m_pie = (this.value & 134217728) != 0;
        return this._m_pie;
      }
    });

    /**
     * Ignore default lib search path.
     */
    Object.defineProperty(DtFlag1Values.prototype, 'nodeflib', {
      get: function() {
        if (this._m_nodeflib !== undefined)
          return this._m_nodeflib;
        this._m_nodeflib = (this.value & 2048) != 0;
        return this._m_nodeflib;
      }
    });

    return DtFlag1Values;
  })();

  var EndianElf = Elf.EndianElf = (function() {
    function EndianElf(_io, _parent, _root) {
      this._io = _io;
      this._parent = _parent;
      this._root = _root || this;

      this._read();
    }
    EndianElf.prototype._read = function() {
      switch (this._root.endian) {
      case Elf.Endian.LE:
        this._is_le = true;
        break;
      case Elf.Endian.BE:
        this._is_le = false;
        break;
      }

      if (this._is_le === true) {
        this._readLE();
      } else if (this._is_le === false) {
        this._readBE();
      } else {
        throw new KaitaiStream.UndecidedEndiannessError();
      }
    }
    EndianElf.prototype._readLE = function() {
      this.eType = this._io.readU2le();
      this.machine = this._io.readU2le();
      this.eVersion = this._io.readU4le();
      switch (this._root.bits) {
      case Elf.Bits.B32:
        this.entryPoint = this._io.readU4le();
        break;
      case Elf.Bits.B64:
        this.entryPoint = this._io.readU8le();
        break;
      }
      switch (this._root.bits) {
      case Elf.Bits.B32:
        this.programHeaderOffset = this._io.readU4le();
        break;
      case Elf.Bits.B64:
        this.programHeaderOffset = this._io.readU8le();
        break;
      }
      switch (this._root.bits) {
      case Elf.Bits.B32:
        this.sectionHeaderOffset = this._io.readU4le();
        break;
      case Elf.Bits.B64:
        this.sectionHeaderOffset = this._io.readU8le();
        break;
      }
      this.flags = this._io.readBytes(4);
      this.eEhsize = this._io.readU2le();
      this.programHeaderEntrySize = this._io.readU2le();
      this.qtyProgramHeader = this._io.readU2le();
      this.sectionHeaderEntrySize = this._io.readU2le();
      this.qtySectionHeader = this._io.readU2le();
      this.sectionNamesIdx = this._io.readU2le();
    }
    EndianElf.prototype._readBE = function() {
      this.eType = this._io.readU2be();
      this.machine = this._io.readU2be();
      this.eVersion = this._io.readU4be();
      switch (this._root.bits) {
      case Elf.Bits.B32:
        this.entryPoint = this._io.readU4be();
        break;
      case Elf.Bits.B64:
        this.entryPoint = this._io.readU8be();
        break;
      }
      switch (this._root.bits) {
      case Elf.Bits.B32:
        this.programHeaderOffset = this._io.readU4be();
        break;
      case Elf.Bits.B64:
        this.programHeaderOffset = this._io.readU8be();
        break;
      }
      switch (this._root.bits) {
      case Elf.Bits.B32:
        this.sectionHeaderOffset = this._io.readU4be();
        break;
      case Elf.Bits.B64:
        this.sectionHeaderOffset = this._io.readU8be();
        break;
      }
      this.flags = this._io.readBytes(4);
      this.eEhsize = this._io.readU2be();
      this.programHeaderEntrySize = this._io.readU2be();
      this.qtyProgramHeader = this._io.readU2be();
      this.sectionHeaderEntrySize = this._io.readU2be();
      this.qtySectionHeader = this._io.readU2be();
      this.sectionNamesIdx = this._io.readU2be();
    }

    var DynsymSectionEntry64 = EndianElf.DynsymSectionEntry64 = (function() {
      function DynsymSectionEntry64(_io, _parent, _root, _is_le) {
        this._io = _io;
        this._parent = _parent;
        this._root = _root || this;
        this._is_le = _is_le;

        this._read();
      }
      DynsymSectionEntry64.prototype._read = function() {

        if (this._is_le === true) {
          this._readLE();
        } else if (this._is_le === false) {
          this._readBE();
        } else {
          throw new KaitaiStream.UndecidedEndiannessError();
        }
      }
      DynsymSectionEntry64.prototype._readLE = function() {
        this.nameOffset = this._io.readU4le();
        this.info = this._io.readU1();
        this.other = this._io.readU1();
        this.shndx = this._io.readU2le();
        this.value = this._io.readU8le();
        this.size = this._io.readU8le();
      }
      DynsymSectionEntry64.prototype._readBE = function() {
        this.nameOffset = this._io.readU4be();
        this.info = this._io.readU1();
        this.other = this._io.readU1();
        this.shndx = this._io.readU2be();
        this.value = this._io.readU8be();
        this.size = this._io.readU8be();
      }

      return DynsymSectionEntry64;
    })();

    var ProgramHeader = EndianElf.ProgramHeader = (function() {
      function ProgramHeader(_io, _parent, _root, _is_le) {
        this._io = _io;
        this._parent = _parent;
        this._root = _root || this;
        this._is_le = _is_le;

        this._read();
      }
      ProgramHeader.prototype._read = function() {

        if (this._is_le === true) {
          this._readLE();
        } else if (this._is_le === false) {
          this._readBE();
        } else {
          throw new KaitaiStream.UndecidedEndiannessError();
        }
      }
      ProgramHeader.prototype._readLE = function() {
        this.type = this._io.readU4le();
        if (this._root.bits == Elf.Bits.B64) {
          this.flags64 = this._io.readU4le();
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.offset = this._io.readU4le();
          break;
        case Elf.Bits.B64:
          this.offset = this._io.readU8le();
          break;
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.vaddr = this._io.readU4le();
          break;
        case Elf.Bits.B64:
          this.vaddr = this._io.readU8le();
          break;
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.paddr = this._io.readU4le();
          break;
        case Elf.Bits.B64:
          this.paddr = this._io.readU8le();
          break;
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.filesz = this._io.readU4le();
          break;
        case Elf.Bits.B64:
          this.filesz = this._io.readU8le();
          break;
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.memsz = this._io.readU4le();
          break;
        case Elf.Bits.B64:
          this.memsz = this._io.readU8le();
          break;
        }
        if (this._root.bits == Elf.Bits.B32) {
          this.flags32 = this._io.readU4le();
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.align = this._io.readU4le();
          break;
        case Elf.Bits.B64:
          this.align = this._io.readU8le();
          break;
        }
      }
      ProgramHeader.prototype._readBE = function() {
        this.type = this._io.readU4be();
        if (this._root.bits == Elf.Bits.B64) {
          this.flags64 = this._io.readU4be();
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.offset = this._io.readU4be();
          break;
        case Elf.Bits.B64:
          this.offset = this._io.readU8be();
          break;
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.vaddr = this._io.readU4be();
          break;
        case Elf.Bits.B64:
          this.vaddr = this._io.readU8be();
          break;
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.paddr = this._io.readU4be();
          break;
        case Elf.Bits.B64:
          this.paddr = this._io.readU8be();
          break;
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.filesz = this._io.readU4be();
          break;
        case Elf.Bits.B64:
          this.filesz = this._io.readU8be();
          break;
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.memsz = this._io.readU4be();
          break;
        case Elf.Bits.B64:
          this.memsz = this._io.readU8be();
          break;
        }
        if (this._root.bits == Elf.Bits.B32) {
          this.flags32 = this._io.readU4be();
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.align = this._io.readU4be();
          break;
        case Elf.Bits.B64:
          this.align = this._io.readU8be();
          break;
        }
      }
      Object.defineProperty(ProgramHeader.prototype, 'dynamic', {
        get: function() {
          if (this._m_dynamic !== undefined)
            return this._m_dynamic;
          if (this.type == Elf.PhType.DYNAMIC) {
            var io = this._root._io;
            var _pos = io.pos;
            io.seek(this.offset);
            if (this._is_le) {
              this._raw__m_dynamic = io.readBytes(this.filesz);
              var _io__raw__m_dynamic = new KaitaiStream(this._raw__m_dynamic);
              this._m_dynamic = new DynamicSection(_io__raw__m_dynamic, this, this._root, this._is_le);
            } else {
              this._raw__m_dynamic = io.readBytes(this.filesz);
              var _io__raw__m_dynamic = new KaitaiStream(this._raw__m_dynamic);
              this._m_dynamic = new DynamicSection(_io__raw__m_dynamic, this, this._root, this._is_le);
            }
            io.seek(_pos);
          }
          return this._m_dynamic;
        }
      });
      Object.defineProperty(ProgramHeader.prototype, 'flagsObj', {
        get: function() {
          if (this._m_flagsObj !== undefined)
            return this._m_flagsObj;
          if (this._is_le) {
            this._m_flagsObj = new PhdrTypeFlags(this._io, this, this._root, (this.flags64 | this.flags32));
          } else {
            this._m_flagsObj = new PhdrTypeFlags(this._io, this, this._root, (this.flags64 | this.flags32));
          }
          return this._m_flagsObj;
        }
      });

      return ProgramHeader;
    })();

    var DynamicSectionEntry = EndianElf.DynamicSectionEntry = (function() {
      function DynamicSectionEntry(_io, _parent, _root, _is_le) {
        this._io = _io;
        this._parent = _parent;
        this._root = _root || this;
        this._is_le = _is_le;

        this._read();
      }
      DynamicSectionEntry.prototype._read = function() {

        if (this._is_le === true) {
          this._readLE();
        } else if (this._is_le === false) {
          this._readBE();
        } else {
          throw new KaitaiStream.UndecidedEndiannessError();
        }
      }
      DynamicSectionEntry.prototype._readLE = function() {
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.tag = this._io.readU4le();
          break;
        case Elf.Bits.B64:
          this.tag = this._io.readU8le();
          break;
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.valueOrPtr = this._io.readU4le();
          break;
        case Elf.Bits.B64:
          this.valueOrPtr = this._io.readU8le();
          break;
        }
      }
      DynamicSectionEntry.prototype._readBE = function() {
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.tag = this._io.readU4be();
          break;
        case Elf.Bits.B64:
          this.tag = this._io.readU8be();
          break;
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.valueOrPtr = this._io.readU4be();
          break;
        case Elf.Bits.B64:
          this.valueOrPtr = this._io.readU8be();
          break;
        }
      }
      Object.defineProperty(DynamicSectionEntry.prototype, 'tagEnum', {
        get: function() {
          if (this._m_tagEnum !== undefined)
            return this._m_tagEnum;
          this._m_tagEnum = this.tag;
          return this._m_tagEnum;
        }
      });
      Object.defineProperty(DynamicSectionEntry.prototype, 'flag1Values', {
        get: function() {
          if (this._m_flag1Values !== undefined)
            return this._m_flag1Values;
          if (this.tagEnum == Elf.DynamicArrayTags.FLAGS_1) {
            if (this._is_le) {
              this._m_flag1Values = new DtFlag1Values(this._io, this, this._root, this.valueOrPtr);
            } else {
              this._m_flag1Values = new DtFlag1Values(this._io, this, this._root, this.valueOrPtr);
            }
          }
          return this._m_flag1Values;
        }
      });

      return DynamicSectionEntry;
    })();

    var SectionHeader = EndianElf.SectionHeader = (function() {
      function SectionHeader(_io, _parent, _root, _is_le) {
        this._io = _io;
        this._parent = _parent;
        this._root = _root || this;
        this._is_le = _is_le;

        this._read();
      }
      SectionHeader.prototype._read = function() {

        if (this._is_le === true) {
          this._readLE();
        } else if (this._is_le === false) {
          this._readBE();
        } else {
          throw new KaitaiStream.UndecidedEndiannessError();
        }
      }
      SectionHeader.prototype._readLE = function() {
        this.ofsName = this._io.readU4le();
        this.type = this._io.readU4le();
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.flags = this._io.readU4le();
          break;
        case Elf.Bits.B64:
          this.flags = this._io.readU8le();
          break;
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.addr = this._io.readU4le();
          break;
        case Elf.Bits.B64:
          this.addr = this._io.readU8le();
          break;
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.ofsBody = this._io.readU4le();
          break;
        case Elf.Bits.B64:
          this.ofsBody = this._io.readU8le();
          break;
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.lenBody = this._io.readU4le();
          break;
        case Elf.Bits.B64:
          this.lenBody = this._io.readU8le();
          break;
        }
        this.linkedSectionIdx = this._io.readU4le();
        this.info = this._io.readBytes(4);
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.align = this._io.readU4le();
          break;
        case Elf.Bits.B64:
          this.align = this._io.readU8le();
          break;
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.entrySize = this._io.readU4le();
          break;
        case Elf.Bits.B64:
          this.entrySize = this._io.readU8le();
          break;
        }
      }
      SectionHeader.prototype._readBE = function() {
        this.ofsName = this._io.readU4be();
        this.type = this._io.readU4be();
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.flags = this._io.readU4be();
          break;
        case Elf.Bits.B64:
          this.flags = this._io.readU8be();
          break;
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.addr = this._io.readU4be();
          break;
        case Elf.Bits.B64:
          this.addr = this._io.readU8be();
          break;
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.ofsBody = this._io.readU4be();
          break;
        case Elf.Bits.B64:
          this.ofsBody = this._io.readU8be();
          break;
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.lenBody = this._io.readU4be();
          break;
        case Elf.Bits.B64:
          this.lenBody = this._io.readU8be();
          break;
        }
        this.linkedSectionIdx = this._io.readU4be();
        this.info = this._io.readBytes(4);
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.align = this._io.readU4be();
          break;
        case Elf.Bits.B64:
          this.align = this._io.readU8be();
          break;
        }
        switch (this._root.bits) {
        case Elf.Bits.B32:
          this.entrySize = this._io.readU4be();
          break;
        case Elf.Bits.B64:
          this.entrySize = this._io.readU8be();
          break;
        }
      }
      Object.defineProperty(SectionHeader.prototype, 'body', {
        get: function() {
          if (this._m_body !== undefined)
            return this._m_body;
          var io = this._root._io;
          var _pos = io.pos;
          io.seek(this.ofsBody);
          if (this._is_le) {
            switch (this.type) {
            case Elf.ShType.DYNAMIC:
              this._raw__m_body = io.readBytes(this.lenBody);
              var _io__raw__m_body = new KaitaiStream(this._raw__m_body);
              this._m_body = new DynamicSection(_io__raw__m_body, this, this._root, this._is_le);
              break;
            case Elf.ShType.STRTAB:
              this._raw__m_body = io.readBytes(this.lenBody);
              var _io__raw__m_body = new KaitaiStream(this._raw__m_body);
              this._m_body = new StringsStruct(_io__raw__m_body, this, this._root, this._is_le);
              break;
            case Elf.ShType.DYNSTR:
              this._raw__m_body = io.readBytes(this.lenBody);
              var _io__raw__m_body = new KaitaiStream(this._raw__m_body);
              this._m_body = new StringsStruct(_io__raw__m_body, this, this._root, this._is_le);
              break;
            case Elf.ShType.DYNSYM:
              this._raw__m_body = io.readBytes(this.lenBody);
              var _io__raw__m_body = new KaitaiStream(this._raw__m_body);
              this._m_body = new DynsymSection(_io__raw__m_body, this, this._root, this._is_le);
              break;
            default:
              this._m_body = io.readBytes(this.lenBody);
              break;
            }
          } else {
            switch (this.type) {
            case Elf.ShType.DYNAMIC:
              this._raw__m_body = io.readBytes(this.lenBody);
              var _io__raw__m_body = new KaitaiStream(this._raw__m_body);
              this._m_body = new DynamicSection(_io__raw__m_body, this, this._root, this._is_le);
              break;
            case Elf.ShType.STRTAB:
              this._raw__m_body = io.readBytes(this.lenBody);
              var _io__raw__m_body = new KaitaiStream(this._raw__m_body);
              this._m_body = new StringsStruct(_io__raw__m_body, this, this._root, this._is_le);
              break;
            case Elf.ShType.DYNSTR:
              this._raw__m_body = io.readBytes(this.lenBody);
              var _io__raw__m_body = new KaitaiStream(this._raw__m_body);
              this._m_body = new StringsStruct(_io__raw__m_body, this, this._root, this._is_le);
              break;
            case Elf.ShType.DYNSYM:
              this._raw__m_body = io.readBytes(this.lenBody);
              var _io__raw__m_body = new KaitaiStream(this._raw__m_body);
              this._m_body = new DynsymSection(_io__raw__m_body, this, this._root, this._is_le);
              break;
            default:
              this._m_body = io.readBytes(this.lenBody);
              break;
            }
          }
          io.seek(_pos);
          return this._m_body;
        }
      });
      Object.defineProperty(SectionHeader.prototype, 'name', {
        get: function() {
          if (this._m_name !== undefined)
            return this._m_name;
          var io = this._root.header.strings._io;
          var _pos = io.pos;
          io.seek(this.ofsName);
          if (this._is_le) {
            this._m_name = KaitaiStream.bytesToStr(io.readBytesTerm(0, false, true, true), "ASCII");
          } else {
            this._m_name = KaitaiStream.bytesToStr(io.readBytesTerm(0, false, true, true), "ASCII");
          }
          io.seek(_pos);
          return this._m_name;
        }
      });
      Object.defineProperty(SectionHeader.prototype, 'flagsObj', {
        get: function() {
          if (this._m_flagsObj !== undefined)
            return this._m_flagsObj;
          if (this._is_le) {
            this._m_flagsObj = new SectionHeaderFlags(this._io, this, this._root, this.flags);
          } else {
            this._m_flagsObj = new SectionHeaderFlags(this._io, this, this._root, this.flags);
          }
          return this._m_flagsObj;
        }
      });

      return SectionHeader;
    })();

    var DynamicSection = EndianElf.DynamicSection = (function() {
      function DynamicSection(_io, _parent, _root, _is_le) {
        this._io = _io;
        this._parent = _parent;
        this._root = _root || this;
        this._is_le = _is_le;

        this._read();
      }
      DynamicSection.prototype._read = function() {

        if (this._is_le === true) {
          this._readLE();
        } else if (this._is_le === false) {
          this._readBE();
        } else {
          throw new KaitaiStream.UndecidedEndiannessError();
        }
      }
      DynamicSection.prototype._readLE = function() {
        this.entries = [];
        var i = 0;
        while (!this._io.isEof()) {
          this.entries.push(new DynamicSectionEntry(this._io, this, this._root, this._is_le));
          i++;
        }
      }
      DynamicSection.prototype._readBE = function() {
        this.entries = [];
        var i = 0;
        while (!this._io.isEof()) {
          this.entries.push(new DynamicSectionEntry(this._io, this, this._root, this._is_le));
          i++;
        }
      }

      return DynamicSection;
    })();

    var DynsymSection = EndianElf.DynsymSection = (function() {
      function DynsymSection(_io, _parent, _root, _is_le) {
        this._io = _io;
        this._parent = _parent;
        this._root = _root || this;
        this._is_le = _is_le;

        this._read();
      }
      DynsymSection.prototype._read = function() {

        if (this._is_le === true) {
          this._readLE();
        } else if (this._is_le === false) {
          this._readBE();
        } else {
          throw new KaitaiStream.UndecidedEndiannessError();
        }
      }
      DynsymSection.prototype._readLE = function() {
        this.entries = [];
        var i = 0;
        while (!this._io.isEof()) {
          switch (this._root.bits) {
          case Elf.Bits.B32:
            this.entries.push(new DynsymSectionEntry32(this._io, this, this._root, this._is_le));
            break;
          case Elf.Bits.B64:
            this.entries.push(new DynsymSectionEntry64(this._io, this, this._root, this._is_le));
            break;
          }
          i++;
        }
      }
      DynsymSection.prototype._readBE = function() {
        this.entries = [];
        var i = 0;
        while (!this._io.isEof()) {
          switch (this._root.bits) {
          case Elf.Bits.B32:
            this.entries.push(new DynsymSectionEntry32(this._io, this, this._root, this._is_le));
            break;
          case Elf.Bits.B64:
            this.entries.push(new DynsymSectionEntry64(this._io, this, this._root, this._is_le));
            break;
          }
          i++;
        }
      }

      return DynsymSection;
    })();

    var DynsymSectionEntry32 = EndianElf.DynsymSectionEntry32 = (function() {
      function DynsymSectionEntry32(_io, _parent, _root, _is_le) {
        this._io = _io;
        this._parent = _parent;
        this._root = _root || this;
        this._is_le = _is_le;

        this._read();
      }
      DynsymSectionEntry32.prototype._read = function() {

        if (this._is_le === true) {
          this._readLE();
        } else if (this._is_le === false) {
          this._readBE();
        } else {
          throw new KaitaiStream.UndecidedEndiannessError();
        }
      }
      DynsymSectionEntry32.prototype._readLE = function() {
        this.nameOffset = this._io.readU4le();
        this.value = this._io.readU4le();
        this.size = this._io.readU4le();
        this.info = this._io.readU1();
        this.other = this._io.readU1();
        this.shndx = this._io.readU2le();
      }
      DynsymSectionEntry32.prototype._readBE = function() {
        this.nameOffset = this._io.readU4be();
        this.value = this._io.readU4be();
        this.size = this._io.readU4be();
        this.info = this._io.readU1();
        this.other = this._io.readU1();
        this.shndx = this._io.readU2be();
      }

      return DynsymSectionEntry32;
    })();

    var StringsStruct = EndianElf.StringsStruct = (function() {
      function StringsStruct(_io, _parent, _root, _is_le) {
        this._io = _io;
        this._parent = _parent;
        this._root = _root || this;
        this._is_le = _is_le;

        this._read();
      }
      StringsStruct.prototype._read = function() {

        if (this._is_le === true) {
          this._readLE();
        } else if (this._is_le === false) {
          this._readBE();
        } else {
          throw new KaitaiStream.UndecidedEndiannessError();
        }
      }
      StringsStruct.prototype._readLE = function() {
        this.entries = [];
        var i = 0;
        while (!this._io.isEof()) {
          this.entries.push(KaitaiStream.bytesToStr(this._io.readBytesTerm(0, false, true, true), "ASCII"));
          i++;
        }
      }
      StringsStruct.prototype._readBE = function() {
        this.entries = [];
        var i = 0;
        while (!this._io.isEof()) {
          this.entries.push(KaitaiStream.bytesToStr(this._io.readBytesTerm(0, false, true, true), "ASCII"));
          i++;
        }
      }

      return StringsStruct;
    })();
    Object.defineProperty(EndianElf.prototype, 'programHeaders', {
      get: function() {
        if (this._m_programHeaders !== undefined)
          return this._m_programHeaders;
        var _pos = this._io.pos;
        this._io.seek(this.programHeaderOffset);
        if (this._is_le) {
          this._raw__m_programHeaders = new Array(this.qtyProgramHeader);
          this._m_programHeaders = new Array(this.qtyProgramHeader);
          for (var i = 0; i < this.qtyProgramHeader; i++) {
            this._raw__m_programHeaders[i] = this._io.readBytes(this.programHeaderEntrySize);
            var _io__raw__m_programHeaders = new KaitaiStream(this._raw__m_programHeaders[i]);
            this._m_programHeaders[i] = new ProgramHeader(_io__raw__m_programHeaders, this, this._root, this._is_le);
          }
        } else {
          this._raw__m_programHeaders = new Array(this.qtyProgramHeader);
          this._m_programHeaders = new Array(this.qtyProgramHeader);
          for (var i = 0; i < this.qtyProgramHeader; i++) {
            this._raw__m_programHeaders[i] = this._io.readBytes(this.programHeaderEntrySize);
            var _io__raw__m_programHeaders = new KaitaiStream(this._raw__m_programHeaders[i]);
            this._m_programHeaders[i] = new ProgramHeader(_io__raw__m_programHeaders, this, this._root, this._is_le);
          }
        }
        this._io.seek(_pos);
        return this._m_programHeaders;
      }
    });
    Object.defineProperty(EndianElf.prototype, 'sectionHeaders', {
      get: function() {
        if (this._m_sectionHeaders !== undefined)
          return this._m_sectionHeaders;
        var _pos = this._io.pos;
        this._io.seek(this.sectionHeaderOffset);
        if (this._is_le) {
          this._raw__m_sectionHeaders = new Array(this.qtySectionHeader);
          this._m_sectionHeaders = new Array(this.qtySectionHeader);
          for (var i = 0; i < this.qtySectionHeader; i++) {
            this._raw__m_sectionHeaders[i] = this._io.readBytes(this.sectionHeaderEntrySize);
            var _io__raw__m_sectionHeaders = new KaitaiStream(this._raw__m_sectionHeaders[i]);
            this._m_sectionHeaders[i] = new SectionHeader(_io__raw__m_sectionHeaders, this, this._root, this._is_le);
          }
        } else {
          this._raw__m_sectionHeaders = new Array(this.qtySectionHeader);
          this._m_sectionHeaders = new Array(this.qtySectionHeader);
          for (var i = 0; i < this.qtySectionHeader; i++) {
            this._raw__m_sectionHeaders[i] = this._io.readBytes(this.sectionHeaderEntrySize);
            var _io__raw__m_sectionHeaders = new KaitaiStream(this._raw__m_sectionHeaders[i]);
            this._m_sectionHeaders[i] = new SectionHeader(_io__raw__m_sectionHeaders, this, this._root, this._is_le);
          }
        }
        this._io.seek(_pos);
        return this._m_sectionHeaders;
      }
    });
    Object.defineProperty(EndianElf.prototype, 'strings', {
      get: function() {
        if (this._m_strings !== undefined)
          return this._m_strings;
        var _pos = this._io.pos;
        this._io.seek(this.sectionHeaders[this.sectionNamesIdx].ofsBody);
        if (this._is_le) {
          this._raw__m_strings = this._io.readBytes(this.sectionHeaders[this.sectionNamesIdx].lenBody);
          var _io__raw__m_strings = new KaitaiStream(this._raw__m_strings);
          this._m_strings = new StringsStruct(_io__raw__m_strings, this, this._root, this._is_le);
        } else {
          this._raw__m_strings = this._io.readBytes(this.sectionHeaders[this.sectionNamesIdx].lenBody);
          var _io__raw__m_strings = new KaitaiStream(this._raw__m_strings);
          this._m_strings = new StringsStruct(_io__raw__m_strings, this, this._root, this._is_le);
        }
        this._io.seek(_pos);
        return this._m_strings;
      }
    });

    return EndianElf;
  })();

  /**
   * File identification, must be 0x7f + "ELF".
   */

  /**
   * File class: designates target machine word size (32 or 64
   * bits). The size of many integer fields in this format will
   * depend on this setting.
   */

  /**
   * Endianness used for all integers.
   */

  /**
   * ELF header version.
   */

  /**
   * Specifies which OS- and ABI-related extensions will be used
   * in this ELF file.
   */

  /**
   * Version of ABI targeted by this ELF file. Interpretation
   * depends on `abi` attribute.
   */

  return Elf;
})();
return Elf;
}));
