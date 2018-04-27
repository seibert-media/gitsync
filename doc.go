// Copyright 2018 The gitsync authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
gitsync provides a service acting as middleware to sync the provided git repository when it's webhook gets called and then calls the next provided webhook while passing all return results back to the original caller
*/

package gitsync
