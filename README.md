# ValiWork - validation framework

[![Build Status](https://ci.dev.pztrn.name/api/badges/pztrn/valiwork/status.svg)](https://ci.dev.pztrn.name/pztrn/valiwork) [![Discord](https://img.shields.io/discord/632359730089689128)](https://discord.gg/CvUnEpM) ![Keybase XLM](https://img.shields.io/keybase/xlm/pztrn)

ValiWork is a validation framework that provides sane API and ability to write own validators that returns arbitrary things. It is goroutine-safe and fast.

## Default validators

There are no necessity to enable default validators at all. But if you want to - call:

```go
valiwork.InitializeDefaultValidators()
```

Default validators will return ``error``.

*There are no default validators ATM. Feel free to submit PR with them!*

## Validators registering and namespacing

Default validators using "T_N" scheme, where ``T`` is data type (string, int, int64, etc.) and ``N`` is a validator name (which can be a generic string). Please, use same naming scheme. Example good validators names:

* ``string_check_for_very_rare_symbol_that_is_not_allowed``
* ``int64_check_if_in_bad_range``
* ``interface_check_if_able_to_be_TheVeryGoodStruct``

Key idea is to help you debugging this thing (see [debug section](#Debug) below).

## Debug

Define ``VALIWORK_DEBUG`` environment variable and set it to ``true`` to get debug output. Default ``log`` module will be used for that.
