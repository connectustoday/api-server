import { prop, Typegoose, ModelType, InstanceType } from 'typegoose';

export default class IUserProfile {
    // @ts-ignore
    @prop({required: true})
	schema_version: string;

	@prop()
	interests?: Array<string>;

	@prop()
	biography?: string;

	@prop()
	education?: string; //TODO

	@prop()
	quote?: string;

	@prop()
	current_residence?: string;

	@prop()
	certifications?: string; //TODO
}
