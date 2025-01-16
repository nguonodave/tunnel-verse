#### 001

A room will never start with the letter `L` or with `#` and must have no spaces.

|          |                                                                                                                                             |
|----------|---------------------------------------------------------------------------------------------------------------------------------------------|
| Problem  | the program goes ahead and uses the bad room name                                                                                           |
| Expected | display an error and exit, or fail to recognise the name so that it fails later with the lengths of connected rooms and room names mismatch |

Run the following command to replicate:

```shell
go run . example04_naming_room_name_should_not_start_with_a_hash
```

```shell
go run . example04_naming_room_name_should_not_start_with_an_L
```

#### 001 (extension)

Since the room names are only restricted from starting with `L`, `space`, and `#`, 
then they can start with any other character. For example, the program, however, does not allow rooms that start with 
the letter `-`:

```shell
go run . example07_room_name_starts_with_valid_letter
```

> [!NOTE]  
> In the above example, the name of `dinish` has been renamed to `-dinish`


#### 002

The input file parser doesn't support comments. This is adverse with the following consequences:
   
   1. If comments appear immediately after the line with the room start marker (`##start`),
      then an error occurs because the parser thinks it's a start room definition.
   
   ```shell
   go run . example04_comment_after_start_hash_tag
   ```
   
   2. Yet again, if comments appear immediately after the line with the room end marker (`##end`),
      then an error occurs because the parser thinks it's an end room definition.
   
   ```shell
   go run . example04_comment_after_end_hash_tag
   ```
   
   3. Comments appearing at the start of the file are flagged as invalid number of ants;
      the parser assumes the number of ants is always in the very first line in the file. Following the file format,
      we should ensure the first valid integer line before the start hashtag is the number of ants, this gives users
      the flexibility of defining comments and leaving spaces at the start of the file.
   
   ```shell
   go run . example04_comment_at_start_of_file
   ```
   
   4. In contrast to points `1` and `2` above, comments with spaces cause errors when placed anywhere
   
   ```shell
   go run . example04_comment_with_spaces
   ```
   
   ```shell
   go run . example04_another_comment_with_spaces
   ```

Expected Behaviour: Comments should be ignored

#### 003

Global variables, as from the `vars` package, make it difficult to test submodules;
need to make submodules independent of global variables.

#### 004

The implementation of the function [FindPaths](processpaths/findpaths.go:19) does not contain any backtracking,
its purely recursion. Need to update the doc comments.






