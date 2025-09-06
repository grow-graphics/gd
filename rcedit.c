/*
 Copyright (c) 2013 GitHub Inc.

 Permission is hereby granted, free of charge, to any person obtaining
 a copy of this software and associated documentation files (the
 "Software"), to deal in the Software without restriction, including
 without limitation the rights to use, copy, modify, merge, publish,
 distribute, sublicense, and/or sell copies of the Software, and to
 permit persons to whom the Software is furnished to do so, subject to
 the following conditions:

 The above copyright notice and this permission notice shall be
 included in all copies or substantial portions of the Software.

 THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
 LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
 OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
 WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>
#include <fcntl.h>
#include <unistd.h>

#define MAX_PATH_COSMO 256
#define MAX_COMMAND 1024

// Helper function to print errors
int print_error(const char* message) {
    fprintf(stderr, "Fatal error: %s\n", message);
    return 1;
}

// Helper function to print warnings
int print_warning(const char* message) {
    fprintf(stderr, "Warning: %s\n", message);
    return 1;
}

// Parse version string (e.g., "1.2.3.4" or shorter)
int parse_version_string(const char* str, unsigned short *v1, unsigned short *v2, unsigned short *v3, unsigned short *v4) {
    *v1 = *v2 = *v3 = *v4 = 0;
    return (sscanf(str, "%hu.%hu.%hu.%hu", v1, v2, v3, v4) == 4) ||
           (sscanf(str, "%hu.%hu.%hu", v1, v2, v3) == 3) ||
           (sscanf(str, "%hu.%hu", v1, v2) == 2) ||
           (sscanf(str, "%hu", v1) == 1);
}

// Generate an .rc file for version info
int generate_version_rc(const char* filename, const char* key, const char* value, const char* file_version, const char* product_version) {
    FILE* rc_file = fopen("temp.rc", "w");
    if (!rc_file) {
        return print_error("Unable to create temporary .rc file");
    }

    fprintf(rc_file,
        "1 VERSIONINFO\n"
        " FILEVERSION %s\n"
        " PRODUCTVERSION %s\n"
        " FILEFLAGSMASK 0x3f\n"
        " FILEFLAGS 0x0\n"
        " FILEOS 0x40004\n" // VOS_NT_WINDOWS32
        " FILETYPE 0x1\n"   // VFT_APP
        " FILESUBTYPE 0x0\n"
        "BEGIN\n"
        "    BLOCK \"StringFileInfo\"\n"
        "    BEGIN\n"
        "        BLOCK \"040904b0\"\n" // US English, Unicode
        "        BEGIN\n"
        "            VALUE \"%s\", \"%s\"\n"
        "        END\n"
        "    END\n"
        "    BLOCK \"VarFileInfo\"\n"
        "    BEGIN\n"
        "        VALUE \"Translation\", 0x409, 1200\n"
        "    END\n"
        "END\n",
        file_version ? file_version : "0,0,0,0",
        product_version ? product_version : "0,0,0,0",
        key, value);

    fclose(rc_file);
    return 0;
}

// Generate an .rc file for icon
int generate_icon_rc(const char* icon_path) {
    FILE* rc_file = fopen("temp.rc", "w");
    if (!rc_file) {
        return print_error("Unable to create temporary .rc file");
    }

    fprintf(rc_file, "IDI_ICON1 ICON DISCARDABLE \"%s\"\n", icon_path);
    fclose(rc_file);
    return 0;
}

// Generate an .rc file for manifest
int generate_manifest_rc(const char* manifest_path) {
    FILE* rc_file = fopen("temp.rc", "w");
    if (!rc_file) {
        return print_error("Unable to create temporary .rc file");
    }

    fprintf(rc_file, "1 RT_MANIFEST \"%s\"\n", manifest_path);
    fclose(rc_file);
    return 0;
}

// Generate an .rc file for RCDATA
int generate_rcdata_rc(const char* key, const char* path) {
    FILE* rc_file = fopen("temp.rc", "w");
    if (!rc_file) {
        return print_error("Unable to create temporary .rc file");
    }

    fprintf(rc_file, "%s RCDATA \"%s\"\n", key, path);
    fclose(rc_file);
    return 0;
}

// Execute zig rc to compile .rc to .res
int compile_rc() {
    char command[MAX_COMMAND] = "./zig rc temp.rc -o temp.res";
    int result = system(command);
    if (result != 0) {
        return print_error("Failed to compile .rc file with zig rc");
    }
    return 0;
}

// Link .res file into executable
int link_resource(const char* target_exe) {
    char command[MAX_COMMAND];
    snprintf(command, sizeof(command), "./zig link -r temp.res %s -o %s", target_exe, target_exe);
    int result = system(command);
    if (result != 0) {
        return print_error("Failed to link resources into executable");
    }
    return 0;
}

// Placeholder for getting version string (not fully implemented due to no PE parsing)
const char* get_version_string(const char* target_exe, const char* key) {
    // TODO: Implement PE parsing to extract version strings
    print_warning("Version string extraction not implemented in this port");
    return NULL;
}

// Placeholder for getting resource string
const char* get_resource_string(const char* target_exe, unsigned int key_id) {
    // TODO: Implement PE parsing to extract resource strings
    print_warning("Resource string extraction not implemented in this port");
    return NULL;
}

int main(int argc, char* argv[]) {
    if (argc == 1 || (argc == 2 && (strcmp(argv[1], "-h") == 0 || strcmp(argv[1], "--help") == 0))) {
        // Hardcode version since we can't use GetFileVersionInfoW
        fprintf(stdout,
            "Rcedit v1.0.0: Edit resources of exe.\n\n"
            "Usage: rcedit <filename> [options...]\n\n"
            "Options:\n"
            "  -h, --help                                 Show this message\n"
            "  --set-version-string <key> <value>         Set version string\n"
            "  --get-version-string <key>                 Print version string\n"
            "  --set-file-version <version>               Set FileVersion\n"
            "  --set-product-version <version>            Set ProductVersion\n"
            "  --set-icon <path-to-icon>                  Set file icon\n"
            "  --set-requested-execution-level <level>    Pass nothing to see usage\n"
            "  --application-manifest <path-to-file>      Set manifest file\n"
            "  --set-resource-string <key> <value>        Set resource string\n"
            "  --get-resource-string <key>                Get resource string\n"
            "  --set-rcdata <key> <path-to-file>          Replace RCDATA by integer id\n");
        return 0;
    }

    int loaded = 0;
    const char* target_exe = NULL;
    const char* file_version = NULL;
    const char* product_version = NULL;
    int has_manifest = 0;

    for (int i = 1; i < argc; ++i) {
        if (strcmp(argv[i], "--set-version-string") == 0 || strcmp(argv[i], "-svs") == 0) {
            if (argc - i < 3) return print_error("--set-version-string requires 'Key' and 'Value'");
            const char* key = argv[++i];
            const char* value = argv[++i];
            if (generate_version_rc(target_exe, key, value, file_version, product_version)) return 1;
            if (compile_rc()) return 1;
            if (link_resource(target_exe)) return 1;

        } else if (strcmp(argv[i], "--get-version-string") == 0 || strcmp(argv[i], "-gvs") == 0) {
            if (argc - i < 2) return print_error("--get-version-string requires 'Key'");
            const char* key = argv[++i];
            const char* result = get_version_string(target_exe, key);
            if (!result) return print_error("Unable to get version string");
            printf("%s\n", result);
            return 0;

        } else if (strcmp(argv[i], "--set-file-version") == 0 || strcmp(argv[i], "-sfv") == 0) {
            if (argc - i < 2) return print_error("--set-file-version requires a version string");
            file_version = argv[++i];
            unsigned short v1, v2, v3, v4;
            if (!parse_version_string(file_version, &v1, &v2, &v3, &v4)) {
                return print_error("Unable to parse version string for FileVersion");
            }
            if (generate_version_rc(target_exe, "FileVersion", file_version, file_version, product_version)) return 1;
            if (compile_rc()) return 1;
            if (link_resource(target_exe)) return 1;

        } else if (strcmp(argv[i], "--set-product-version") == 0 || strcmp(argv[i], "-spv") == 0) {
            if (argc - i < 2) return print_error("--set-product-version requires a version string");
            product_version = argv[++i];
            unsigned short v1, v2, v3, v4;
            if (!parse_version_string(product_version, &v1, &v2, &v3, &v4)) {
                return print_error("Unable to parse version string for ProductVersion");
            }
            if (generate_version_rc(target_exe, "ProductVersion", product_version, file_version, product_version)) return 1;
            if (compile_rc()) return 1;
            if (link_resource(target_exe)) return 1;

        } else if (strcmp(argv[i], "--set-icon") == 0 || strcmp(argv[i], "-si") == 0) {
            if (argc - i < 2) return print_error("--set-icon requires path to the icon");
            const char* icon_path = argv[++i];
            if (generate_icon_rc(icon_path)) return 1;
            if (compile_rc()) return 1;
            if (link_resource(target_exe)) return 1;

        } else if (strcmp(argv[i], "--set-requested-execution-level") == 0 || strcmp(argv[i], "-srel") == 0) {
            if (argc - i < 2) return print_error("--set-requested-execution-level requires asInvoker, highestAvailable or requireAdministrator");
            if (has_manifest) {
                print_warning("--set-requested-execution-level is ignored if --application-manifest is set");
                ++i;
                continue;
            }
            const char* level = argv[++i];
            // Generate a manifest file with the requested execution level
            FILE* manifest = fopen("temp_manifest.xml", "w");
            if (!manifest) return print_error("Unable to create temporary manifest file");
            fprintf(manifest,
                "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n"
                "<assembly xmlns=\"urn:schemas-microsoft-com:asm.v1\" manifestVersion=\"1.0\">\n"
                "  <trustInfo xmlns=\"urn:schemas-microsoft-com:asm.v3\">\n"
                "    <security>\n"
                "      <requestedPrivileges>\n"
                "        <requestedExecutionLevel level=\"%s\" uiAccess=\"false\"/>\n"
                "      </requestedPrivileges>\n"
                "    </security>\n"
                "  </trustInfo>\n"
                "</assembly>\n", level);
            fclose(manifest);
            if (generate_manifest_rc("temp_manifest.xml")) return 1;
            if (compile_rc()) return 1;
            if (link_resource(target_exe)) return 1;

        } else if (strcmp(argv[i], "--application-manifest") == 0 || strcmp(argv[i], "-am") == 0) {
            if (argc - i < 2) return print_error("--application-manifest requires local path");
            has_manifest = 1;
            const char* manifest_path = argv[++i];
            if (generate_manifest_rc(manifest_path)) return 1;
            if (compile_rc()) return 1;
            if (link_resource(target_exe)) return 1;

        } else if (strcmp(argv[i], "--set-resource-string") == 0 || strcmp(argv[i], "-srs") == 0) {
            if (argc - i < 3) return print_error("--set-resource-string requires int 'Key' and string 'Value'");
            const char* key = argv[++i];
            const char* value = argv[++i];
            if (generate_rcdata_rc(key, value)) return 1;
            if (compile_rc()) return 1;
            if (link_resource(target_exe)) return 1;

        } else if (strcmp(argv[i], "--get-resource-string") == 0 || strcmp(argv[i], "-grs") == 0) {
            if (argc - i < 2) return print_error("--get-resource-string requires int 'Key'");
            const char* key = argv[++i];
            unsigned int key_id = atoi(key);
            const char* result = get_resource_string(target_exe, key_id);
            if (!result) return print_error("Unable to get resource string");
            printf("%s\n", result);
            return 0;

        } else if (strcmp(argv[i], "--set-rcdata") == 0) {
            if (argc - i < 3) return print_error("--set-rcdata requires int 'Key' and path to resource 'Value'");
            const char* key = argv[++i];
            const char* path = argv[++i];
            if (generate_rcdata_rc(key, path)) return 1;
            if (compile_rc()) return 1;
            if (link_resource(target_exe)) return 1;

        } else {
            if (loaded) {
                fprintf(stderr, "Unrecognized argument: \"%s\"\n", argv[i]);
                return 1;
            }
            loaded = 1;
            target_exe = argv[i];
            if (access(target_exe, F_OK) != 0) {
                fprintf(stderr, "Unable to load file: \"%s\"\n", target_exe);
                return 1;
            }
        }
    }

    if (!loaded) return print_error("You should specify an exe/dll file");

    // Clean up temporary files
    unlink("temp.rc");
    unlink("temp.res");
    unlink("temp_manifest.xml");

    return 0;
}
