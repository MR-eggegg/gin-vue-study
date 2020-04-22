TEXT ·_benchmark(SB), $0-8

    MOVQ buf+0(FP), DI

    WORD $0x3145; BYTE $0xdb       // xor    r11d, r11d
    LONG $0x03e8b841; WORD $0x0000 // mov    r8d, 1000
    WORD $0xd231                   // xor    edx, edx

LBB0_1:
    WORD $0x894d; BYTE $0xd9 // mov    r9, r11
    LONG $0xc9af0f4d         // imul    r9, r9
    WORD $0x894d; BYTE $0xc2 // mov    r10, r8
    WORD $0x3145; BYTE $0xf6 // xor    r14d, r14d

LBB0_2:
    WORD $0x894c; BYTE $0xf6 // mov    rsi, r14
    LONG $0xf6af0f48         // imul    rsi, rsi
    WORD $0x014c; BYTE $0xce // add    rsi, r9
    WORD $0x894c; BYTE $0xd0 // mov    rax, r10
    WORD $0xc931             // xor    ecx, ecx

LBB0_3:
    WORD $0x8948; BYTE $0xcb     // mov    rbx, rcx
    LONG $0xdbaf0f48             // imul    rbx, rbx
    WORD $0x3948; BYTE $0xde     // cmp    rsi, rbx
    JNE  LBB0_6
    WORD $0x8548; BYTE $0xc0     // test    rax, rax
    JNE  LBB0_6
    LONG $0xd71c8b48             // mov    rbx, qword [rdi + 8*rdx]
    WORD $0x894c; BYTE $0x1b     // mov    qword [rbx], r11
    LONG $0xd75c8b48; BYTE $0x08 // mov    rbx, qword [rdi + 8*rdx + 8]
    WORD $0x894c; BYTE $0x33     // mov    qword [rbx], r14
    LONG $0xd75c8b48; BYTE $0x10 // mov    rbx, qword [rdi + 8*rdx + 16]
    WORD $0x8948; BYTE $0x0b     // mov    qword [rbx], rcx
    LONG $0x03c28348             // add    rdx, 3

LBB0_6:
    WORD $0xff48; BYTE $0xc1                   // inc    rcx
    WORD $0xff48; BYTE $0xc8                   // dec    rax
    LONG $0xe9f98148; WORD $0x0003; BYTE $0x00 // cmp    rcx, 1001
    JNE  LBB0_3
    WORD $0xff49; BYTE $0xc6                   // inc    r14
    WORD $0xff49; BYTE $0xca                   // dec    r10
    LONG $0xe9fe8149; WORD $0x0003; BYTE $0x00 // cmp    r14, 1001
    JNE  LBB0_2
    WORD $0xff49; BYTE $0xc3                   // inc    r11
    WORD $0xff49; BYTE $0xc8                   // dec    r8
    LONG $0xe9fb8149; WORD $0x0003; BYTE $0x00 // cmp    r11, 1001
    JNE  LBB0_1
    RET
